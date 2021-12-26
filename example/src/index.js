const go = new Go();
WebAssembly.instantiateStreaming(fetch('timeconverter.wasm'), go.importObject).then((result) => {
  go.run(result.instance);
  initApp();
});

function initApp() {
  initHours();
}

function initHours() {
  const inputHours = document.getElementById('input-hours');
  const outputHours = document.getElementById('output-hours');
  function inputHoursToOutput() {
    const hoursToSeconds = timeconverter.hoursToSeconds(Number(inputHours.value));
    outputHours.value = hoursToSeconds;
  }
  inputHoursToOutput();
  inputHours.addEventListener('input', inputHoursToOutput);
}
