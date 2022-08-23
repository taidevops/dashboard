import {Injectable} from '@angular/core';

@Injectable({providedIn: 'root'})
export class DecoderService {
  base64(encoded: string): string {
    try {
      return decodeURIComponent(
        window
          .atob(encoded)
          .split('')
          .map((c: string) => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
          .join('')
      );
    } catch (e) {
      return 'Unable to decode file. It might contain binary data.';
    }
  }
}
