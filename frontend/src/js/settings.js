import { BASE_URL } from './config.js';
const newTableForm = $('#new-table-form');
const newRowForm = $('#new-row-form');
const dropdownCodeList = $('#code-list');
const dropdownMenuButton = $('#dropdownMenuButton');
let newTableBody = null;
let tableName = null;

$(function () {
  main();
});

function main() {
  dropdownMenuButton.on('click', () => {
    dropdownCodeList.html('');

    const settings = {
      url: `${BASE_URL}/credit/codes_list`,
      method: 'GET',
      contentType: 'application/json',
      dataType: 'json',
      timeout: 0,
    };

    $.ajax(settings)
      .done(response => {
        response.forEach(el => {
          const newCodeItem = $(createCodeItem(el.Name, el.Code));

          newCodeItem.on('click', () => fetchDataByCode(el));

          dropdownCodeList.append(newCodeItem);
        });
      })
      .fail(error => {
        console.error(`Error Info
        Status: ${error.status}
        Text: ${error.statusText}`);
      });
  });

  newTableForm.on('submit', e => {
    e.preventDefault();
    tableName = $('#new-table').val().trim();
    $('#table-name-title').text(tableName);
    createNewTable();
    newTableForm.css('display', 'none');
    newRowForm.css('display', 'grid');

    const settings = {
      url: `${BASE_URL}/credit/create_table`,
      method: 'POST',
      data: JSON.stringify({ db_name: tableName }),
      contentType: 'application/json',
      dataType: 'json',
      timeout: 0,
    };

    $.ajax(settings)
      .done(response => {})
      .fail(error => {
        console.error(`Error Info
          Status: ${error.status}
          Text: ${error.statusText}`);
      });
  });

  newRowForm.on('submit', e => {
    e.preventDefault();
    newTableBody = $('#new-table-body');
    const newRowDateElement = $('#new-row-date');
    const newRowPriceElement = $('#new-row-price');
    let newRowDate = newRowDateElement.val();
    let newRowPrice = +newRowPriceElement.val();

    const newTableRow = createTableRow(newRowDate, newRowPrice);

    newTableBody.append(newTableRow);

    newRowDateElement.val('');
    newRowPriceElement.val('');
  });
}

function createCodeItem(name, code) {
  return `<li><span class="dropdown-item c-pointer" data-code="${code}">${name}</span></li>`;
}

function createNewTable() {
  const newTableElement = `<table
    id="new-table"
    class="table table-striped table-hover table-responsive">
        <thead>
            <tr>
                <th scope="col">Дата</th>
                <th scope="col">Средняя цена</th>
                <th scope="col"></th>
            </tr>
        </thead>
        <tbody id="new-table-body"></tbody>
    </table>`;

  $('#table-box').html(newTableElement);
}

function createTableRow(date, midprice) {
  const tableRow = `<tr>
      <td>${date}</td>
      <td>
        <input
        type="number"
        step=".01"
        class="form-control"
        placeholder="Средняя цена"
        value="${midprice.toFixed(2)}" />
      </td>
      <td>
        <button class="delete-icon text-danger">
            <i class="fa-solid fa-trash"></i>
        </button>
      </td>
    </tr>`;

  return tableRow;
}

function fetchDataByCode(el) {
  const settings = {
    url: `${BASE_URL}/credit/${el.Code}/retrieve_two_columns/`,
    method: 'GET',
    contentType: 'application/json',
    dataType: 'json',
    timeout: 0,
  };

  $.ajax(settings)
    .done(response => {
      console.log(el.Name);
      $('#table-name-title').text(el.Name);

      const yearsData = response.dataset.data;

      createNewTable();

      for (let [key, value] of Object.entries(yearsData)) {
        let year = value[0];
        let price = value[1];

        const newTableRow = createTableRow(year, price);

        $('#new-table-body').append(newTableRow);
      }
    })
    .fail(error => {
      console.error(`Error Info
            Status: ${error.status}
            Text: ${error.statusText}`);
    });
}
