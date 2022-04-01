import { NgModule } from '@angular/core';
import { GongngxchartsspecificComponent } from './gongngxchartsspecific.component';
import { GongngxchartsChartComponent } from './gongngxcharts-chart/gongngxcharts-chart.component';



@NgModule({
  declarations: [
    GongngxchartsspecificComponent,
    GongngxchartsChartComponent
  ],
  imports: [
  ],
  exports: [
    GongngxchartsspecificComponent,
    GongngxchartsChartComponent
  ]
})
export class GongngxchartsspecificModule { }
