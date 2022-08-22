import { Input } from "@angular/core";

export class CardComponent {
  @Input() initialized = true;
  @Input() role: 'inner' | 'table' | 'inner-content';
  
}