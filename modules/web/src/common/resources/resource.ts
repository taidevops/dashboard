import {HttpClient} from '@angular/common/http';

export abstract class ResourceBase {
  protected constructor(protected readonly http_: HttpClient) {}
}