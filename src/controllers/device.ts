import { Request, Response } from 'express'
import { Prisma, PrismaClient } from '@prisma/client'


const prisma = new PrismaClient()

export async function create(req: Request, res: Response) {
  try {
    const { id, secret, name, enabled, zoneId } = req.body;

    if (!id)
      return res.status(400).json({ error: "Missing device id" })

    if (!secret)
      return res.status(400).json({ error: "Missing device secret" })

    if (!name)
      return res.status(400).json({ error: "Missing device name" })

    if (!zoneId)
      return res.status(400).json({ error: "Missing device zone" })



    const new_device = await prisma.device.create({
      data: { id, name, secret, enabled, zoneId }
    })

    res.status(201).json(new_device)
  } catch (error) {
    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      return res.status(500).json({ error: error.meta?.cause })
    } else {
      console.error(error)
      res.status(500).json({ error: "Internal Server Error" })
    }
  }
}

export async function edit(req: Request, res: Response) {
  try {
    const { id } = req.params
    const { secret, name, enabled, zoneId } = req.body;

    if (!id)
      return res.status(400).json({ error: "Missing device id" })

    const device = await prisma.device.update({
      where: { id },
      data: { name, secret, enabled, zoneId }
    })

    res.status(200).json(device)
  } catch (error) {
    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      return res.status(500).json({ error: error.meta?.cause })
    } else {
      console.error(error)
      res.status(500).json({ error: "Internal Server Error" })
    }
  }
}

export async function get(req: Request, res: Response) {
  try {
    const { id } = req.params

    if (!id)
      return res.status(404).json({ error: "Invalid device id" })


    const device = await prisma.device.findUniqueOrThrow({
      where: { id },
      include: { zone: true }
    })

    res.status(200).json(device)
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

export async function remove(req: Request, res: Response) {
  try {
    const { id } = req.params

    if (!id)
      return res.status(404).json({ error: "Invalid device id" })

    const device = await prisma.device.delete({
      where: { id },
    })

    res.status(200).json(device)
  } catch (error) {
    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      return res.status(500).json({ error: error.meta?.cause })
    } else {
      console.error(error)
      res.status(500).json({ error: "Internal Server Error" })
    }
  }
}

export async function list(req: Request, res: Response) {
  try {
    const page = Number(req.query.page) || 1
    const page_size = Number(req.query.size) || 10

    const devices = await prisma.device.findMany({
      orderBy: { name: 'asc' },
      skip: (page - 1) * page_size,
      take: page_size
    })

    res.status(200).json(devices)
  } catch (error) {
    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      return res.status(500).json({ error: error.meta?.cause })
    } else {
      console.error(error)
      res.status(500).json({ error: "Internal Server Error" })
    }
  }
}

