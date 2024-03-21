import Link from "next/link";
import * as React from "react";
import { Button } from "../form";

function HeaderComponent() {
  return (
    <header className="sticky top-0 bg-white">
      <nav className="hidden md:flex justify-between items-center container py-6">
        <div className="flex items-center gap-4">
          <Link href={"/"} className="">
            <p>Captioner</p>
          </Link>
          <Link href={"#"} className="">
            <p className="font-medium">Resources</p>
          </Link>
          <Link href={"/pricing"} className="">
            <p className="font-medium">Pricing</p>
          </Link>
        </div>
        <div className="">
          <Link href={"/auth/signin"} className="mr-4">
            <Button variant={"ghost"}>Login</Button>
          </Link>
          <Link href={"/auth/signup"}>
            <Button size={"lg"}>Get Started</Button>
          </Link>
        </div>
      </nav>
      <nav className="flex md:hidden justify-between items-center container py-6">
        <Link href={"/"} className="text-base font-bold">
          <p>Captioner</p>
        </Link>
        <div>
          <div className=" items-center gap-4 hidden">
            <Link href={"#"} className="">
              <p className="font-medium">Resources</p>
            </Link>
            <Link href={"/pricing"} className="">
              <p className="font-medium">Pricing</p>
            </Link>
          </div>
          <div className="hidden">
            <Link href={"/auth/signin"} className="mr-4">
              <Button variant={"ghost"}>Login</Button>
            </Link>
            <Link href={"/auth/signup"}>
              <Button size={"lg"}>Get Started</Button>
            </Link>
          </div>
        </div>
      </nav>
      <hr />
    </header>
  );
}

export default HeaderComponent;
