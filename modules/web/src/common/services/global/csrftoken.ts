import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {CsrfToken} from '@api/root.api';
import {Observable} from 'rxjs';

@Injectable()
export class CsrfTokenService {
  constructor(private readonly http_: HttpClient) {}

  /** Get a CSRF token for an action you want to perform. */
  getTokenForAction(action: string): Observable<CsrfToken> {
    return this.http_.get<CsrfToken>(`api/v1/csrftoken/${action}`);
  }
}
