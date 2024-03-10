"use client";
import * as React from "react";
import {
  FacebookIcon,
  InstagramIcon,
  LinkedInIcon,
  TiktokIcon,
  TwitterIcon,
} from "../icons";
import { animate, motion, useMotionValue } from "framer-motion";

function HomepagePlatforms() {
  const x = useMotionValue(0);
  React.useEffect(() => {
    animate(x, -200, {
      ease: "linear",
      duration: 5,
      repeat: Infinity,
      delay: 0,
    });
  }, []);
  return (
    <section className="container py-4">
      <div className="w-[192px] mx-auto overflow-hidden">
        <motion.div className="flex items-center gap-2" style={{ x }}>
          <div className=" w-8 h--8">
            <FacebookIcon width={32} height={32} />
          </div>
          <div className=" w-8 h--8">
            <InstagramIcon width={32} height={32} />
          </div>
          <div className=" w-8 h--8">
            <LinkedInIcon width={32} height={32} />
          </div>
          <div className=" w-8 h--8">
            <TiktokIcon width={32} height={32} />
          </div>
          <div className=" w-8 h--8">
            <TwitterIcon width={32} height={32} />
          </div>

          <div className=" w-8 h--8">
            <FacebookIcon width={32} height={32} />
          </div>
          <div className=" w-8 h--8">
            <InstagramIcon width={32} height={32} />
          </div>
          <div className=" w-8 h--8">
            <LinkedInIcon width={32} height={32} />
          </div>
          <div className=" w-8 h--8">
            <TiktokIcon width={32} height={32} />
          </div>
          <div className=" w-8 h--8">
            <TwitterIcon width={32} height={32} />
          </div>
        </motion.div>
      </div>
    </section>
  );
}

export default HomepagePlatforms;
