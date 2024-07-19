import { Request, Response } from 'express'
import { Prisma, PrismaClient } from '@prisma/client'


const prisma = new PrismaClient()

export async function create(req: Request, res: Response) {
  try {
    const { name, code } = req.body;

    if (!name)
      return res.status(400).json({ error: "Missing zone name" })

    if (!code)
      return res.status(400).json({ error: "Missing zone code" })

    const new_zone = await prisma.zone.create({
      data: { name, code }
    })

    res.status(201).json(new_zone)
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
    const { name, code } = req.body;

    if (!id)
      return res.status(400).json({ error: "Invalid zone id" })


    const zone = await prisma.zone.update({
      where: { id: Number(id!) },
      data: {
        name,
        code
      }
    })

    res.status(200).json(zone)
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

    if (!Number(id))
      return res.status(400).json({ error: "Invalid zone id" })


    const zone = await prisma.zone.findUniqueOrThrow({
      where: { id: Number(id!) },
      include: {
        devices: true
      }
    })

    res.status(200).json(zone)
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
      return res.status(400).json({ error: "Invalid zone id" })

    const zone = await prisma.zone.delete({
      where: { id: Number(id!) },
    })

    res.status(200).json(zone)
  } catch (error) {
    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      return res.status(400).json({ error: error.meta.cause })
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

    const zones = await prisma.zone.findMany({
      orderBy: { name: 'asc' },
      skip: (page - 1) * page_size,
      take: page_size
    })

    res.status(200).json(zones)
  } catch (error) {
    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      return res.status(500).json({ error: error.meta?.cause })
    } else {
      console.error(error)
      res.status(500).json({ error: "Internal Server Error" })
    }
  }
}

