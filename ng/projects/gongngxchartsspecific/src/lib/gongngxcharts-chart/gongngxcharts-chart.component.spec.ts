import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongngxchartsChartComponent } from './gongngxcharts-chart.component';

describe('GongngxchartsChartComponent', () => {
  let component: GongngxchartsChartComponent;
  let fixture: ComponentFixture<GongngxchartsChartComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongngxchartsChartComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(GongngxchartsChartComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
