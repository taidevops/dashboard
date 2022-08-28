import { EventEmitter, Injectable } from "@angular/core";
import { Theme } from "@api/root.api";
import { ThemeSwitchCallback } from "@api/root.ui";

@Injectable()
export class ThemeService {
  static readonly SystemTheme = '__system_theme__';
  private _customThemes: Theme[] = [];
  private readonly _defaultThemes: Theme[] = [
    {name: 'kd-light-theme', displayName: 'Light', isDark: false},
    {name: 'kd-dark-theme', displayName: 'Dark', isDark: true},
  ];
  private readonly _onThemeSwitchEvent = new EventEmitter<string>();
  constructor(

  ) {}

  private _theme = ThemeService.SystemTheme;

  get theme(): string {
    return this._theme;
  }

  subscribe(callback: ThemeSwitchCallback): void {
    this._onThemeSwitchEvent.subscribe(callback);
  }
}