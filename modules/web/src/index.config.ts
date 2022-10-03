import {InjectionToken} from '@angular/core';
import {MatTooltipDefaultOptions} from '@angular/material/tooltip';
import { IConfig } from '@api/root.ui';

export const CONFIG_DI_TOKEN = new InjectionToken<IConfig>('kd.config');

// Override default material tooltip values.
export const KD_TOOLTIP_DEFAULT_OPTIONS: MatTooltipDefaultOptions = {
  showDelay: 500,
  hideDelay: 0,
  touchendHideDelay: 0,
};