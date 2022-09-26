import { APP_INITIALIZER, NgModule } from "@angular/core";
import {DecoderService} from '@common/services/global/decoder';
import { LocalConfigLoaderService } from "./loader";

@NgModule({
  providers: [
    DecoderService,
    {
      provide: APP_INITIALIZER,
      useFactory: init,
      deps: [
        LocalConfigLoaderService,
      ]
    }
  ]
})
export class GlobalServicesModule {

}

export function init(
  loader: LocalConfigLoaderService
): Function {
  return () => {
    return loader.init().then(() => {

    });
  };
}