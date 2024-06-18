export interface GetUpdatesBody {
  force: boolean;
  onlyInstalledVersion?: boolean;
}

export interface VersionInfo {
  fullVersion?: string;
  timestamp?: string;
  version?: string;
  tag?: string;
}

export interface GetUpdatesResponse {
  lastCheck: string;
  latest?: VersionInfo;
  installed: VersionInfo;
  latestNewsUrl?: string;
  updateAvailable?: boolean;
}

export interface StartUpdateBody {
  newImage?: string;
}

export interface StartUpdateResponse {
  authToken: string;
  logOffset: number;
}

export interface GetUpdateStatusBody {
  authToken: string;
  logOffset: number;
}

export interface GetUpdateStatusResponse {
  done: boolean;
  logOffset: number;
  logLines: string[];
}