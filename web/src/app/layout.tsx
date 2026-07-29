import './globals.css'
import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'GabraOS — Engineering Control Center',
  description: 'The Open Standard & Operating System for Autonomous Engineering',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className="antialiased min-h-screen bg-[#080c14] text-slate-100">{children}</body>
    </html>
  )
}
