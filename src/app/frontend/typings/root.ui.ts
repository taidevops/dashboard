
export type ThemeSwitchCallback = (isLightThemeEnabled: boolean) => void;

export interface AppConfig {

}

export enum IMessageKey {
  Open = 'Open',
  Close = 'Close',
  Pin = 'Pin'
}

export type IMessage = {
  [key in IMessageKey]: string;
};
