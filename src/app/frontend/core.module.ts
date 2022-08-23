import {Inject, NgModule, Optional, SkipSelf} from '@angular/core';
import {CookieService} from 'ngx-cookie-service';

import {GlobalServicesModule} from '@common/services/global/module';
import {ResourceModule} from '@common/services/resource/module';

@NgModule({
  providers: [
    CookieService,
  ],
  imports: [GlobalServicesModule, ResourceModule]
})
export class CoreModule {
  /* make sure CoreModule is imported only by one NgModule the RootModule */
  constructor(@Inject(CoreModule) @Optional() @SkipSelf() parentModule: CoreModule) {
    if (parentModule) {
      throw new Error('CoreModule is already loaded. Import only in RootModule.');
    }
  }
}