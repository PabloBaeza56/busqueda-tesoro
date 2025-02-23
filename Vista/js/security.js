window.onload = function () {
  // Evitar que el usuario retroceda o avance en el navegador
  window.history.pushState(null, null, window.location.href);

  window.onpopstate = function (event) {
    window.history.go(1);
  };

  // Evitar que el usuario utilice el click derecho
  document.addEventListener("contextmenu", function (e) {
    e.preventDefault();
  });
};

//Desactivar el botón de entrar después de oprimirlo
function desactivarBotonDeEnvio() {
  const boton = document.getElementById('botonEnvio');
  boton.disabled = true;
  boton.textContent = 'Procesando...';
}