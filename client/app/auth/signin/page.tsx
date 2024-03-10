import { Button } from "@/components/form";
import * as React from "react";

export default function Signin(): React.JSX.Element {
  return (
    <>
      <div className="max-w-screen-sm w-full">
        <p>Captioner</p>
        <p>Signin</p>

        <div className="border rounded-lg p-10 w-full">
          <form className="flex flex-col">
            <label htmlFor="email">
              <p>Email</p>
              <input type="email" name="" id="" />
            </label>
            <Button className="w-full">Continue</Button>
          </form>
          <p>or</p>
          <div>
            <Button size={'lg'} className="w-full">Continue with Google</Button>
          </div>
        </div>
      </div>
    </>
  );
}
