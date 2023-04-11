// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: April 2023
// This file contains the JS functions for index.html

'use strict'

/**
 * This function calculates the volume of a sphere
 */

function myButtonClicked() {
  // input
  const radius = parseFloat(document.getElementById('radius').value)

  // process
  const volume = (4 / 3) * Math.PI * radius ** 3

  // output
  document.getElementById('volume').innerHTML = "Volume is " + volume.toFixed(2) + " cm³"
}