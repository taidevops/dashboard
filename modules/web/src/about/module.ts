import { NgModule } from "@angular/core";
import { AboutComponent } from "./component";
import { AboutRoutingModule } from "./routing";

@NgModule({
  imports: [AboutRoutingModule],
  declarations: [AboutComponent]
})
export class AboutModule {}