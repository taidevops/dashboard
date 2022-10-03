
import {CommonModule} from '@angular/common';
import {NgModule} from '@angular/core';

import {MAT_TOOLTIP_DEFAULT_OPTIONS, MatTooltipModule} from '@angular/material/tooltip';

import {KD_TOOLTIP_DEFAULT_OPTIONS} from './index.config';

const SHARED_DEPENDENCIES = [
  // Angular imports
  CommonModule,

]

@NgModule({
  imports: SHARED_DEPENDENCIES,
  exports: SHARED_DEPENDENCIES,
  providers: [
    {
      provide: MAT_TOOLTIP_DEFAULT_OPTIONS,
      useValue: KD_TOOLTIP_DEFAULT_OPTIONS,
    },
  ],
})
export class SharedModule {}
