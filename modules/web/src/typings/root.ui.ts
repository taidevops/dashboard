export type ThemeSwitchCallback = (isLightThemeEnabled: boolean) => void;

export interface VersionInfo {
  dirty: boolean;
  raw: string;
  hash: string;
  distance: number;
  tag: string;
  semver: SemverInfo;
  suffix: string;
  semverString: string;
  packageVersion: string;
}

export interface SemverInfo {
  raw: string;
  major: number;
  minor: number;
  patch: number;
  prerelease: string[];
  build: string[];
  version: string;
  loose: boolean;
  options: SemverInfoOptions;
}

export interface SemverInfoOptions {
  loose: boolean;
  includePrerelease: boolean;
}

export interface RatioItem {
  name: string;
  value: number;
}

export interface ResourcesRatio {
  cronJobRatio: RatioItem[];
  daemonSetRatio: RatioItem[];
  deploymentRatio: RatioItem[];
  jobRatio: RatioItem[];
  podRatio: RatioItem[];
}

export interface AppConfig {

}

export interface IConfig {
  authTokenCookieName: string;
  usernameCookieName: string;
  skipLoginPageCookieName: string;
  csrfHeaderName: string;
  authTokenHeaderName: string;
  defaultNamespace: string;
  authModeCookieName: string;
  supportedLanguages: LanguageConfig[];
  defaultLanguage: string;
  languageCookieName: string;
}

export interface LanguageConfig {
  value: string;
  label: string;
}

export enum IMessageKey {
  Open = 'Open',
  Close = 'Close',
  Pin = 'Pin'
}

export type IMessage = {
  [key in IMessageKey]: string;
};

export enum IBreadcrumbMessageKey {
  Logs = 'Logs',
  Error = 'Error',
  Create = 'Create',
}

export type IBreadcrumbMessage = {
  [key in IBreadcrumbMessageKey]: string;
};
