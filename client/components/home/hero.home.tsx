import * as React from "react";
import { Button } from "../form";
import Link from "next/link";

function HomepageHero() {
  return (
    <>
      <section>
        <div className="container text-center py-[70px] max-w-screen-lg">
          <div className="mb-8">
            <h1 className="text-4xl md:text-6xl font-bold mb-4 text-[#171717] leading-[1.2] max-w-screen-sm mx-auto">
              Run & manage your <br /> socials efficiently
            </h1>
            <p className="text-[#171717]/60">
              Content Calenders, Schedule Posting, Insight and Analytics for all
              your socials in one place.{" "}
            </p>
          </div>
          <div>
            <Link href={"/auth/signup"}>
              <Button size={"lg"}>Try For Free*</Button>
            </Link>
            <p className="text-xs">
              *Premium features require paid accounts, prices in NGN.
            </p>
          </div>
        </div>
      </section>
    </>
  );
}

export default HomepageHero;
