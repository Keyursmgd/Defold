components {
  id: "enemybullet"
  component: "/main/enemybullet.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/assests/ship.atlas\"\n"
  "default_animation: \"New Animation\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: -0.7051087
    w: 0.70909923
  }
  scale {
    x: 1.0
    y: 0.867612
    z: 1.0
  }
}
