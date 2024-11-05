#!/bin/bash -e
set -o errexit
set -o nounset

usage() {
  cat <<-EOF
Usage: $BASE_NAME [OPTIONS]
Options:
      --platform <platform>    Build for a specific platform (defaults to linux/amd64)
      --no-update              Do not fetch the latest changes from the repo
      --update-only            Only fetch the latest changes from the repo
      --no-client              Do not build PMM Client
      --no-client-docker       Do not build PMM Client docker image (default)
      --client-only            Build only PMM Client (client binaries + docker)
      --no-server              Do not build PMM Server (docker image)
      --log-file <path>        Save build logs to a file located at <path> (defaults to PWD)
      --help | -h              Display help
EOF
}

parse-params() {
  # Define global variables
  LOCAL_BUILD=1
  NO_UPDATE=0
  UPDATE_ONLY=0
  NO_CLIENT=0
  NO_CLIENT_DOCKER=1
  NO_SERVER=0
  START_TIME=$(date +%s)
  LOG_FILE="$(dirname $0)/build.log"
  BASE_NAME=$(basename $0)
  PLATFORM=linux/amd64
  SUBMODULES=pmm-submodules
  PATH_TO_SCRIPTS="sources/pmm/src/github.com/percona/pmm/build/scripts"

  while test "$#" -gt 0; do
    case "$1" in
      --update-only)
        UPDATE_ONLY=1; NO_UPDATE=0
        ;;
      --no-update)
        if [ "$UPDATE_ONLY" -eq 1 ]; then
          echo "Error. Mutually exclusive options: --update-only and --no-update"
          exit 1
        fi      
        NO_UPDATE=1
        ;;
      --client-only)
        NO_CLIENT=0; NO_CLIENT_DOCKER=0; NO_SERVER=1
        ;;
      --no-client)
        NO_CLIENT=1; NO_CLIENT_DOCKER=1
        ;;
      --client-docker)
        NO_CLIENT_DOCKER=0
        ;;
      --no-client-docker)
        if [ "$NO_CLIENT" -eq 1 ]; then
          echo "Error. Mutually exclusive options: --client-docker and --no-client"
          exit 1
        fi
        if [ "$NO_CLIENT_DOCKER" -eq 1 ]; then
          echo "Error. Mutually exclusive options: --client-docker and --no-client-docker"
          exit 1
        fi
        NO_CLIENT_DOCKER=1
        ;;
      --no-server)
        NO_SERVER=1
        ;;
      --platform)
        shift
        if [ -z "$1" ]; then
          echo "Missing argument for --platform"
          exit 1
        fi
        PLATFORM="$1"
        ;;
      --log-file)
        shift
        if [ -z "$1" ]; then
          echo "Missing argument for --log-file"
          exit 1
        fi
        LOG_FILE="$1"
        ;;
      --help | -h)
        shift
        usage
        exit 0
        ;;
      *)
        echo "Unknown argument: $1"
        usage
        exit 1
        ;;
    esac
    shift
  done
}

needs-to-pull() {
  local UPSTREAM=${1:-'@{u}'}
  local LOCAL=$(git rev-parse @)
  local BASE=$(git merge-base @ "$UPSTREAM")
  local REMOTE=$(git rev-parse "$UPSTREAM")

  if [ "$LOCAL" = "$REMOTE" ]; then
    return 1 # false, we are up-to-date
  fi

  if [ "$LOCAL" = "$BASE" ]; then
    return 0 # true, we are behind upstream
  fi
}

rewind() {
  local DIR="$1"
  local BRANCH="$2"

  cd "$DIR" > /dev/null
  local CURRENT=$(git branch --show-current)
  git fetch

  if [ "$CURRENT" != "$BRANCH" ]; then
    echo "Currently on $CURRENT, checking out $BRANCH"
    git checkout "$BRANCH"
  fi

  if needs-to-pull; then
    git pull origin
    echo "Submodule has pulled from upstream"
    git logs -n 2
    cd - > /dev/null
    git add "$DIR"
  else
    cd - > /dev/null
    echo "Submodule is up-to-date with upstream"
  fi
}

check-files() {
  local DIR="$1"

  test -z "DIR" && exit 1

  if [ -d "$DIR/sources" ] && [ -f "$DIR/ci-default.yml" ] && [ -f "$DIR/ci.yml" ]; then
    return 0
  fi

  return 1
}

# Update submodules and PR branches
update() {
  local DEPS=
  local CURDIR="$PWD"
  local UPDATED_SCRIPT="$SUBMODULES/$PATH_TO_SCRIPTS/build/local/build.sh"
  local MD5SUM=$(md5sum $(dirname $0)/build.sh)

   if [ "$NO_UPDATE" -eq 1 ]; then
    echo "Running without refreshing the source code from repositories..."
    return
   fi

  # Thouroughly verify the presence of known files, otherwise bail out
  if [ ! -d "$SUBMODULES" ] ; then # pwd must be outside of pmm-submodules
    echo "Warn: the current working directory must be outside of pmm-submodules"
    echo "cd .."
    cd .. > /dev/null
  fi

  if [ -d "$SUBMODULES" ]; then # pwd is outside pmm-submodules
    if ! check-files "$SUBMODULES"; then
      echo "Fatal: could not locate known files in ${PWD}/${SUBMODULES}"
      exit 1
    fi
  else
    echo "Fatal: could not locate known files in $PWD"
    exit 1
  fi

  cd "$SUBMODULES"

  # Join the dependencies from ci-default.yml and ci.yml
  DEPS=$(yq -o=json eval-all '. as $item ireduce ({}; . *d $item )' ci-default.yml ci.yml | jq '.deps')

  echo "This script rewinds submodule branches as per the joint config of 'ci-default.yml' and 'ci.yml'"

  echo "$DEPS" | jq -c '.[]' | while read -r item; do
    branch=$(echo "$item" | jq -r '.branch')
    path=$(echo "$item" | jq -r '.path')
    name=$(echo "$item" | jq -r '.name')
    echo
    echo "Rewinding submodule '$name' ..."
    echo "path: ${path}, branch: ${branch}"

    rewind "$path" "$branch"
  done

  echo
  echo "Printing git status..."
  git status --short
  echo
  echo "Printing git submodule status..."
  git submodule status

  cd "$CURDIR" > /dev/null

  if [ -f "$UPDATED_SCRIPT" ] && [ "$MD5SUM" != "$(md5sum $UPDATED_SCRIPT)" ]; then
    echo "The local copy of this script differs from the one fetched from the repo." 
    echo "Apparently, that version is newer. We will halt to give you the change to run a fresh version."
    echo "You can copy it over and run it again, i.e. '/bin/bash $(dirname $0)/build.sh --no-update'"
    exit 0
  fi
}

get_branch_name() {
  local module="${1:-}"
  local path=$(git config -f .gitmodules submodule.${module}.path)

  if [ ! -d "$path" ]; then
    echo "Error: could not resolve the path to submodule ${module}"
    exit 1
  fi

  echo $(git -C "$path" branch --show-current)
}

run_build_script() {
  local CURDIR="$PWD"
  local script="$PATH_TO_SCRIPTS/$1"
  local script_name="$1"
  local start_time
  local end_time

  cd "$SUBMODULES" > /dev/null

  if [ ! -f "$script" ]; then
    echo "Fatal: script $script does not exist."
    cd "$CURDIR" > /dev/null
    exit 1
  fi

  start_time=$(date +%s)
  if [ "$#" -gt 1 ]; then
    shift
    script_name="${script_name}:($1)"
    $script "$@"
  else
    $script
  fi
  end_time=$(date +%s)

  echo ---
  echo "Execution time (in sec) for $script_name: $((end_time - start_time))"
  echo ---

  cd "$CURDIR" > /dev/null
}

purge_files() {
  local CURDIR=$PWD
  local tmp_files

  cd "$SUBMODULES" > /dev/null
  # Remove stale files and directories
  if [ -d tmp ]; then
    echo "Removing stale files and directories..."

    if [ -d "tmp/pmm-server" ]; then
      tmp_files=$(find tmp/pmm-server | grep -v "RPMS" | grep -Ev "^tmp/pmm-server$" || :)
      if [ -n "$tmp_files" ]; then
        tmp_files=( $tmp_files )
        for f in "${tmp_files[@]}"; do
          echo "Removing file or directory $f ..."
          rm -rf "$f"
        done
      fi
    fi

    if [ -d "tmp/source/pmm" ]; then
      echo "Removing tmp/source/pmm ..."
      rm -rf tmp/source/pmm
    fi
  fi
  
  cd "$CURDIR"
}

init() {
  local CURDIR="$PWD"

  if [ -d "$SUBMODULES" ]; then
    cd "$SUBMODULES" > /dev/null
  fi

  export RPMBUILD_DOCKER_IMAGE=perconalab/rpmbuild:3

  GIT_COMMIT=$(git rev-parse HEAD | head -c 8)

  # Create docker volume to persist package and build cache
  # Read more in the section about `rpmbuild`.
  if ! docker volume ls | grep pmm-gobuild >/dev/null; then
    docker volume create pmm-gobuild
  fi
  if ! docker volume ls | grep pmm-gomod >/dev/null; then
    docker volume create pmm-gomod
  fi
  if ! docker volume ls | grep pmm-yarn >/dev/null; then
    docker volume create pmm-yarn
  fi
  if ! docker volume ls | grep pmm-dnf >/dev/null; then
    docker volume create pmm-dnf
  fi

  cd "$CURDIR" > /dev/null
}

cleanup() {
  local CURDIR="$PWD"
  cd "$SUBMODULES" > /dev/null

  # Implement cleanup logic here

  cd "$CURDIR" > /dev/null
}

main() {
  update

  init

  purge_files

  if [ "$NO_CLIENT" -eq 0 ]; then
    # Build client source: 4m39s from scratch, 0m27s using cache
    run_build_script build-client-source

    # Build client binary: ??? from scratch, 0m20s using cache
    run_build_script build-client-binary

    # Building client source rpm takes 13s (caching does not apply)
    run_build_script build-client-srpm

    # Building client rpm takes 1m40s
    run_build_script build-client-rpm
  fi

  # Building client docker image takes 17s
  export DOCKER_CLIENT_TAG=percona/pmm-client:${GIT_COMMIT}
  if [ "$NO_CLIENT_DOCKER" -eq 0 ]; then
    run_build_script build-client-docker
  fi

  if [ "$NO_SERVER" -eq 0 ]; then
    # Grafana build fails to compile with Go 1.23.x, see https://github.com/grafana/grafana/issues/89796
    # We need to apply this [patch](https://github.com/grafana/grafana/pull/94742) to fix it
    export RPMBUILD_DOCKER_IMAGE=perconalab/rpmbuild:3
    export DOCKER_TAG=percona/pmm-server:${GIT_COMMIT}
    export DOCKERFILE=Dockerfile.el9
    export RPM_EPOCH=1

    # 3rd-party components
    run_build_script build-server-rpm grafana
    run_build_script build-server-rpm victoriametrics

    # 1st-party components
    run_build_script build-server-rpm percona-dashboards grafana-dashboards
    run_build_script build-server-rpm pmm-managed pmm
    run_build_script build-server-rpm percona-qan-api2 pmm
    run_build_script build-server-rpm pmm-update pmm
    run_build_script build-server-rpm pmm-dump
    run_build_script build-server-rpm vmproxy pmm

    run_build_script build-server-docker
  fi

  echo
  echo "Done building PMM artifacts."
  echo ---
  echo "Total execution time, sec: $(($(date +%s) - $START_TIME))"
  echo ---

  cleanup
}

parse-params "$@"

# Capture the build logs in the log file
exec > >(tee "$LOG_FILE") 2>&1

main