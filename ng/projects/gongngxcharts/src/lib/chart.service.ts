// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { ChartDB } from './chart-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class ChartService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ChartServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private chartsUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.chartsUrl = origin + '/api/github.com/fullstack-lang/gongngxcharts/go/v1/charts';
  }

  /** GET charts from the server */
  getCharts(): Observable<ChartDB[]> {
    return this.http.get<ChartDB[]>(this.chartsUrl)
      .pipe(
        tap(_ => this.log('fetched charts')),
        catchError(this.handleError<ChartDB[]>('getCharts', []))
      );
  }

  /** GET chart by id. Will 404 if id not found */
  getChart(id: number): Observable<ChartDB> {
    const url = `${this.chartsUrl}/${id}`;
    return this.http.get<ChartDB>(url).pipe(
      tap(_ => this.log(`fetched chart id=${id}`)),
      catchError(this.handleError<ChartDB>(`getChart id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new chart to the server */
  postChart(chartdb: ChartDB): Observable<ChartDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    return this.http.post<ChartDB>(this.chartsUrl, chartdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`posted chartdb id=${chartdb.ID}`)
      }),
      catchError(this.handleError<ChartDB>('postChart'))
    );
  }

  /** DELETE: delete the chartdb from the server */
  deleteChart(chartdb: ChartDB | number): Observable<ChartDB> {
    const id = typeof chartdb === 'number' ? chartdb : chartdb.ID;
    const url = `${this.chartsUrl}/${id}`;

    return this.http.delete<ChartDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted chartdb id=${id}`)),
      catchError(this.handleError<ChartDB>('deleteChart'))
    );
  }

  /** PUT: update the chartdb on the server */
  updateChart(chartdb: ChartDB): Observable<ChartDB> {
    const id = typeof chartdb === 'number' ? chartdb : chartdb.ID;
    const url = `${this.chartsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    return this.http.put<ChartDB>(url, chartdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated chartdb id=${chartdb.ID}`)
      }),
      catchError(this.handleError<ChartDB>('updateChart'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
