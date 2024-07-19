import { Router } from "express";
import { verifyToken } from './oauth'
import { MqttClient } from "mqtt";
import { onRequestAuthenticate, onRequestOpen } from "./controllers";
import * as device from "./controllers/device";
import * as zone from "./controllers/zone";



export function setup_admin_router(router: Router) {
  router.get("/devices", device.list)
  router.get("/device/:id", device.get)
  router.post("/device", device.create)
  router.post("/device/:id", device.edit)
  router.delete("/device/:id", device.remove)

  router.get("/zones", zone.list)
  router.get("/zone/:id", zone.get)
  router.post("/zone", zone.create)
  router.post("/zone/:id", zone.edit)
  router.delete("/zone/:id", zone.remove)

  return router
}

export function setup_api_router(router: Router, broker?: MqttClient) {

  router.use(verifyToken)

  router.post("/authenticate", onRequestAuthenticate)

  broker && router.post("/open", onRequestOpen(broker))

  return router
}
