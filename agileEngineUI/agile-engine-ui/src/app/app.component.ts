import {Component, OnInit} from '@angular/core';
import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {animate, state, style, transition, trigger} from '@angular/animations';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({height: '0px', minHeight: '0'})),
      state('expanded', style({height: '*'})),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
@Injectable()
export class AppComponent implements OnInit{

  allTransaction: TransactionModel[];
  columnsToDisplay = ['type', 'amount'];
  expandedElement: TransactionModel | null;

  constructor(private httpClient: HttpClient) {
  }

  getAllTransaction() {
    this.httpClient.get<Array<TransactionModel>>('http://localhost:8001/alltransactions').subscribe((response) => {
      this.allTransaction = response;
      console.log(response);
    });

  }

  ngOnInit(): void {
    this.getAllTransaction();
  }


}
