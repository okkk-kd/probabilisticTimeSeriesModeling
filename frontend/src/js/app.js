import { BASE_URL } from './config.js';

$(function () {
  main();
});

function main() {
  const loginTitle = $('#login-title');
  const regTitle = $('#reg-title');
  const loginForm = $('#login-form');
  const regForm = $('#reg-form');

  loginTitle.click(() => {
    loginTitle.addClass('active');
    regTitle.removeClass('active');
    loginForm.delay(300).fadeIn(300);
    regForm.fadeOut(300);
  });

  regTitle.click(() => {
    regTitle.addClass('active');
    loginTitle.removeClass('active');
    regForm.delay(300).fadeIn(300);
    loginForm.fadeOut(300);
  });

  loginForm.on('submit', e => {
    e.preventDefault();
    auth();
  });

  regForm.on('submit', e => {
    e.preventDefault();
    auth('register');
  });
}

function auth(authType = 'login') {
  let user_name = $('#user_name').val().trim();
  let password = $('#password').val();
  let url = `${BASE_URL}/user/authorization`;

  if (authType === 'register') {
    user_name = $('#reg-user_name').val().trim();
    password = $('#reg-password').val();
    url = `${BASE_URL}/user/registration`;
  }

  const settings = {
    url,
    method: 'POST',
    data: JSON.stringify({ user_name, password }),
    contentType: 'application/json',
    dataType: 'json',
    timeout: 0,
  };

  $.ajax(settings)
    .done(response => {
      localStorage.setItem('session_key', response.session_key);

      $(location).attr(
        'href',
        'http://127.0.0.1:5500/src/pages/home/index.html'
      );
    })
    .fail(error => {
      console.error(`Error Info
      Status: ${error.status}
      Text: ${error.statusText}`);

      if (authType === 'login') {
        $('#password').val('');
        $('#error').css('display', 'block');
      } else {
        $('#reg-password').val('');
        $('#reg-error').css('display', 'block');
      }
    });
}
