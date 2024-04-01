%global _dwz_low_mem_die_limit 0

%global repo            pmm
%global provider        github.com/percona/%{repo}
%global commit          8f3d007617941033867aea6a134c48b39142427f
%global shortcommit     %(c=%{commit}; echo ${c:0:7})
%define release         21
%define rpm_release     %{release}.%{shortcommit}%{?dist}

# the line below is sed'ed by build/bin/build-server-rpm to set a correct version
%define full_pmm_version 2.0.0

%if ! 0%{?gobuild:1}
# https://github.com/rpm-software-management/rpm/issues/367
# https://fedoraproject.org/wiki/PackagingDrafts/Go#Build_ID
%define gobuild(o:) go build -ldflags "${LDFLAGS:-} -B 0x$(head -c20 /dev/urandom | od -An -tx1 | tr -d ' \\n')" -a -v -x %{?**};
%endif

Name:     pmm-managed
Version:  %{version}
Release:  %{rpm_release}
Summary:  Percona Monitoring and Management management daemon

License:  AGPLv3
URL:      https://%{provider}
Source0:  https://%{provider}/archive/%{commit}/%{repo}-%{shortcommit}.tar.gz

%description
pmm-managed manages configuration of PMM server components (VictoriaMetrics,
Grafana, etc.) and exposes API for that. Those APIs are used by pmm-admin tool.
See PMM docs for more information.


%prep
%setup -q -n %{repo}-%{commit}

%build
export PMM_RELEASE_VERSION=%{full_pmm_version}
export PMM_RELEASE_FULLCOMMIT=%{commit}
export PMM_RELEASE_BRANCH=""

make -C managed release


%install
install -d -p %{buildroot}%{_bindir}
install -d -p %{buildroot}%{_sbindir}
install -d -p %{buildroot}%{_datadir}/%{name}
install -p -m 0755 bin/pmm-managed %{buildroot}%{_sbindir}/pmm-managed
install -p -m 0755 bin/pmm-managed-init %{buildroot}%{_sbindir}/pmm-managed-init
install -p -m 0755 bin/pmm-managed-starlark %{buildroot}%{_sbindir}/pmm-managed-starlark

cp -pa api/swagger %{buildroot}%{_datadir}/%{name}


%files
%license managed/LICENSE
%doc managed/README.md
%{_sbindir}/pmm-managed
%{_sbindir}/pmm-managed-init
%{_sbindir}/pmm-managed-starlark
%{_datadir}/%{name}


%changelog
* Mon Apr 1 2024 Alex Demidoff <alexander.demidoff@percona.com> - 3.0.0-21
- PMM-12899 Use module and build cache

* Thu Jul 28 2022 Alex Tymchuk <alexander.tymchuk@percona.com> - 2.30.0-1
- PMM-10036 migrate to monorepo

* Fri Jun 17 2022 Anton Bystrov <anton.bystrov@simbirsoft.com> - 2.0.0-17
- PMM-10206 merge pmm-managed to monorepo pmm

* Thu Jul  2 2020 Mykyta Solomko <mykyta.solomko@percona.com> - 2.0.0-17
- PMM-5645 built using Golang 1.14

* Tue May 12 2020 Alexey Palazhchenko <alexey.palazhchenko@percona.com> - 2.0.0-16
- added pmm-managed-starlark

* Tue Feb 11 2020 Mykyta Solomko <mykyta.solomko@percona.com> - 2.0.0-14
- added pmm-managed-init

* Thu Sep  5 2019 Viacheslav Sarzhan <slava.sarzhan@percona.com> - 2.0.0-10
- init version
