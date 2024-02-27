// Setup Apline Store with state machine
document.addEventListener('alpine:init', () => {
  Alpine.store('projectModal', {
    on: false,
    toggle() {
      this.on = !this.on
    },
  })
  Alpine.store('assignUserModal', {
    on: false,
    toggle() {
      this.on = !this.on
    },
  })
  Alpine.store('editProjectModal', {
    on: false,
    id: "",
    toggle(id) {
      this.on = !this.on
      if (this.on) {
        this.id = id
      } else {
        this.id = ""
      }
    },
  })

})

