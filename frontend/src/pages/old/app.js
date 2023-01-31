$(function () {
  main();
});

function main() {
  const settings = {
    url: 'http://localhost:4000/Points',
    method: 'GET',
    timeout: 0,
  };

  $.ajax(settings)
    .done(response => {
      onSuccess(response);
    })
    .fail(error => {
      console.error(`Error Info
    Status: ${error.status}
    Text: ${error.statusText}`);
    });
}

function onSuccess(data) {
  const priceTableBody = $('#price-table-body');

  for (let [key, value] of Object.entries(data)) {
    const rowData = value[0];

    const newTableRow = createTableRow(key, rowData.MidPrice, rowData.Date);

    priceTableBody.append(newTableRow);
  }
}

function createTableRow(year, midprice, date) {
  const tableRow = `<tr>
    <td>${year}</td>
    <td>${midprice.toFixed(2)}</td>
    <td>${date.split('-').reverse().join('.')}</td>
  </tr>`;

  return tableRow;
}
