
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