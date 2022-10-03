
export interface CsrfToken {
  token: string;
}

export interface Theme {
  name: string;
  displayName: string;
  isDark: boolean;
}

export interface AppConfig {
  serverTime: number;
}

export interface EnvVar {
  name: string;
  value: string;
}

export interface Container {
  name: string;
  image: string;
  env: EnvVar[];
  commands: string[];
  args: string[];
}

export interface LoginStatus {
  tokenPresent: boolean;
  headerPresent: boolean;
  httpsMode: boolean;
  impersonationPresent?: boolean;
  impersonatedUser?: string;
}

export type AuthenticationMode = string;

export interface EnabledAuthenticationModes {
  modes: AuthenticationMode[];
}

export interface LoginSkippableResponse {
  skippable: boolean;
}
