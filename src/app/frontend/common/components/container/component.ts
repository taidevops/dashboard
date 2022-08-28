import { Input, OnChanges, SimpleChanges } from "@angular/core";
import { Container } from "@api/root.api";

export class ContainerCardComponent implements OnChanges {
  @Input() container: Container;
  @Input() namespace: string;
  @Input() initialized: boolean;

  constructor(private readonly state_: string) {}

  ngOnChanges(): void {
    this.container.env = this.container.env.sort((a, b) => a.name.localeCompare(b.name));
  }
}