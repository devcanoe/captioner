import Link from "next/link";
import * as React from "react";

function FooterComponent() {
  return (
    <>
      <footer className="container py-[35px]">
        <div className="flex items-center justify-between">
          <ul>
            <p>Resources</p>
            <li>
              <Link href={"#"}>Blog</Link>
            </li>
            <li>
              <Link href={"#"}>F.A.Qs</Link>
            </li>
            <li>
              <Link href={"#"}>Blog</Link>
            </li>
          </ul>
          <ul>
            <p>Resources</p>
            <li>
              <Link href={"#"}>Blog</Link>
            </li>
            <li>
              <Link href={"#"}>F.A.Qs</Link>
            </li>
            <li>
              <Link href={"#"}>Blog</Link>
            </li>
          </ul>
          <ul>
            <p>Resources</p>
            <li>
              <Link href={"#"}>Blog</Link>
            </li>
            <li>
              <Link href={"#"}>F.A.Qs</Link>
            </li>
            <li>
              <Link href={"#"}>Blog</Link>
            </li>
          </ul>
          <ul>
            <p>Resources</p>
            <li>
              <Link href={"#"}>Blog</Link>
            </li>
            <li>
              <Link href={"#"}>F.A.Qs</Link>
            </li>
            <li>
              <Link href={"#"}>Blog</Link>
            </li>
          </ul>
        </div>

        <p className="text-center">Captioner © 2024</p>
      </footer>
    </>
  );
}

export default FooterComponent;
