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

import { PkgeltDB } from './pkgelt-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class PkgeltService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  PkgeltServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private pkgeltsUrl: string

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
    this.pkgeltsUrl = origin + '/api/github.com/fullstack-lang/gongdoc/go/v1/pkgelts';
  }

  /** GET pkgelts from the server */
  getPkgelts(): Observable<PkgeltDB[]> {
    return this.http.get<PkgeltDB[]>(this.pkgeltsUrl)
      .pipe(
        tap(_ => this.log('fetched pkgelts')),
        catchError(this.handleError<PkgeltDB[]>('getPkgelts', []))
      );
  }

  /** GET pkgelt by id. Will 404 if id not found */
  getPkgelt(id: number): Observable<PkgeltDB> {
    const url = `${this.pkgeltsUrl}/${id}`;
    return this.http.get<PkgeltDB>(url).pipe(
      tap(_ => this.log(`fetched pkgelt id=${id}`)),
      catchError(this.handleError<PkgeltDB>(`getPkgelt id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new pkgelt to the server */
  postPkgelt(pkgeltdb: PkgeltDB): Observable<PkgeltDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    pkgeltdb.Classdiagrams = []
    pkgeltdb.Umlscs = []

    return this.http.post<PkgeltDB>(this.pkgeltsUrl, pkgeltdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`posted pkgeltdb id=${pkgeltdb.ID}`)
      }),
      catchError(this.handleError<PkgeltDB>('postPkgelt'))
    );
  }

  /** DELETE: delete the pkgeltdb from the server */
  deletePkgelt(pkgeltdb: PkgeltDB | number): Observable<PkgeltDB> {
    const id = typeof pkgeltdb === 'number' ? pkgeltdb : pkgeltdb.ID;
    const url = `${this.pkgeltsUrl}/${id}`;

    return this.http.delete<PkgeltDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted pkgeltdb id=${id}`)),
      catchError(this.handleError<PkgeltDB>('deletePkgelt'))
    );
  }

  /** PUT: update the pkgeltdb on the server */
  updatePkgelt(pkgeltdb: PkgeltDB): Observable<PkgeltDB> {
    const id = typeof pkgeltdb === 'number' ? pkgeltdb : pkgeltdb.ID;
    const url = `${this.pkgeltsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    pkgeltdb.Classdiagrams = []
    pkgeltdb.Umlscs = []

    return this.http.put<PkgeltDB>(url, pkgeltdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated pkgeltdb id=${pkgeltdb.ID}`)
      }),
      catchError(this.handleError<PkgeltDB>('updatePkgelt'))
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
