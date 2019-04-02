workflow "CI" {
  on = "push"
  resolves = [
    "build"
  ]
}

action "build" {
  uses = "myles-systems/actions-golang@master"
  args = "make test"
}
