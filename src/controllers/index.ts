import { PrismaClient, Prisma } from '@prisma/client';
import { Request, Response } from 'express'
import { MqttClient } from 'mqtt/*';
import { TOTP } from "totp-generator"

const prisma = new PrismaClient()

interface RequestWithUser extends Request {
  user?: {
    iss: string,
    azp: string,
    aud: string,
    sub: string,
    hd: string,
    email: string,
    email_verified: boolean,
    nbf: number,
    name: string,
    picture: string,
    given_name: string,
    family_name: string,
    iat: string,
    exp: string,
    jti: string
  }
}

export function onRequestAuthenticate(req: RequestWithUser, res: Response) {
  const authorized = !!req.user
  res.status(authorized ? 200 : 401).json({ authorized });
}

export function onRequestOpen(broker: MqttClient) {
  return async (req: RequestWithUser, res: Response) => {
    try {
      const { door_id, otp_code } = req.body

      if (!door_id)
        return res.status(404).json({ error: "Invalid device id" })

      const device = await prisma.device.findUniqueOrThrow({ where: { id: door_id } })
      const { otp } = TOTP.generate(device.secret, { period: 60, timestamp: Date.now() / 1000 })

      if (otp_code == otp) {
        broker.publish(door_id, otp_code)
        res.status(200).send()
      } else {
        res.status(401).send()
      }

    } catch (error) {
      if (error instanceof Prisma.NotFoundError) {
        return res.status(404).json({ error: error.message })
      } else if (error instanceof Prisma.PrismaClientKnownRequestError) {
        return res.status(500).json({ error: error.meta?.cause })
      } else {
        console.error(error)
        res.status(500).json({ error: "Internal Server Error" })
      }
    }
  }
} 