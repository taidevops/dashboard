import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {ResourceBase} from '../../resources/resource';

@Injectable()
export class ResourceService<T> extends ResourceBase {
  constructor (readonly http: HttpClient) {
    super(http);
  }

  get(endpoint: string, name?: string, params?: HttpParams): Observable<T> {
    if (name) {
      endpoint = endpoint.replace(':name', name);
    }

    return this
  }
}