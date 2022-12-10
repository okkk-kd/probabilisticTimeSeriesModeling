const BASE_URL = 'http://127.0.0.1:7001/api';

$(function () {
  main();
});

function main() {
  const dropdownCodeList = $('#code-list');

  const settings = {
    url: `${BASE_URL}/credit/bank_forecast`,
    method: 'GET',
    contentType: 'application/json',
    dataType: 'json',
    timeout: 0,
  };

  $.ajax(settings)
    .done(response => {
      response.forEach(el => {
        const newCodeItem = $(createCodeItem(el.Name, el.Code));

        newCodeItem.on('click', () => fetchDataByCode(el.Code));

        dropdownCodeList.append(newCodeItem);
      });
    })
    .fail(error => {
      console.error(`Error Info
      Status: ${error.status}
      Text: ${error.statusText}`);
    });

  //   drawChart();
}

function createCodeItem(name, code) {
  return `<li><span class="dropdown-item c-pointer" data-code="${code}">${name}</span></li>`;
}

function fetchDataByCode(code) {
  const settings = {
    url: `${BASE_URL}/credit/${code}/retrieve_two_columns/`,
    method: 'GET',
    contentType: 'application/json',
    dataType: 'json',
    timeout: 0,
  };

  $.ajax(settings)
    .done(response => {
      const yearsData = response.dataset.data;

      const avgPrices = yearsData.map(year => year[1]).reverse();
      const years = yearsData.map(year => +year[0].split('-')[0]).reverse();

      drawChart(years, avgPrices);
    })
    .fail(error => {
      console.error(`Error Info
          Status: ${error.status}
          Text: ${error.statusText}`);
    });
}

function drawChart(xData, yData) {
  const canvasBox = $('#canvas-box');
  const ctx = $('#myChart');

  if (ctx) {
    ctx.remove();
    const newCtx = $('<canvas id="myChart"></canvas>');
    canvasBox.append(newCtx);

    new Chart(newCtx, {
      type: 'line',
      data: {
        labels: xData,
        datasets: [
          {
            label: 'Средняя цена',
            data: yData,
            borderWidth: 1,
          },
        ],
      },
      options: {
        scales: {
          y: {
            beginAtZero: true,
          },
        },
      },
    });
  }
}
