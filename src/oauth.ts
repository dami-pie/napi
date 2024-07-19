import { Request, Response, NextFunction } from "express";
import { OAuth2Client } from 'google-auth-library'

const oauth2Client = new OAuth2Client();

export async function verifyToken(req: Request, res: Response, next: NextFunction) {
  try {
    const idToken = req.headers.authorization?.split(' ')[0];

    if (!idToken)
      return res.status(401).json({ error: 'Invalid token' })

    const ticket = await oauth2Client.verifyIdToken({
      idToken,
      audience: process.env.OAUTH_CLIENT_ID
    })

    const payload = ticket.getPayload();

    if ((payload.exp * 1000) < Date.now())
      return res.status(401).json({ error: 'Token expired' });

    (req as any).user = payload;
    next();

  } catch (err) {
    return res.status(500).json({ error: 'Internal Server Error' })
  }
}