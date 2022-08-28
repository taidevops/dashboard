import { Component, OnInit } from "@angular/core";
import { ThemeService } from "@common/services/global/theme";

@Component({selector: 'kd-root', template: '<router-outlet></router-outlet>'})
export class RootComponent implements OnInit {
  private _theme = this._themeService.theme;

  constructor(
    private readonly _themeService: ThemeService,
  ) {}

  ngOnInit(): void {
    this._themeService.subscribe(this.onThemeChange_.bind(this));
  }

  private onThemeChange_(theme: string): void {
    this._theme = theme;
  }
}