import { NgModule } from "@angular/core";
import {RouterModule} from '@angular/router';
import {ResourceService} from './resource';

@NgModule({
  imports: [RouterModule],
  providers: [ResourceService]
})
export class ResourceModule {}