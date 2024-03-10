import * as React from "react";

export default function AuthLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>): React.JSX.Element {
  return <main className="grid place-content-center min-h-screen">{children}</main>;
}
