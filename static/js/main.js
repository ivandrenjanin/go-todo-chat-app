// Setup Apline Store with state machine
document.addEventListener("alpine:init", () => {
  Alpine.store("projectModal", {
    on: false,
    toggle() {
      this.on = !this.on;
    },
  });
});
