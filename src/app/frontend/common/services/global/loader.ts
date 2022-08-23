import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {AppConfig} from '@api/root.ui';
import {tap} from 'rxjs/operators';

@Injectable()
export class LocalConfigLoaderService {
  private appConfig_: AppConfig = {} as AppConfig;

  constructor(private readonly http_: HttpClient) {}

  get appConfig(): AppConfig {
    return this.appConfig_;
  }

  init(): Promise<{}> {
    return this.http_
      .get('assets/config/config.json')
      .pipe(tap(response => (this.appConfig_ = response as AppConfig)))
      .toPromise();
  }
}
