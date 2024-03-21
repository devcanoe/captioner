import * as React from "react";

function HomepageAbout() {
  return (
    <section className="bg-black rounded md:py-[140px] py-[35px]">
      <div className="container px-4">
        <h2 className="text-3xl md:text-5xl text-white font-medium mb-10">
          Automate Growth
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div className="w-full text-white flex items-start gap-2">
            <p className="text-5xl">1</p>
            <div>
              <p className="mb-2 text-lg font-medium">Content Calenders</p>
              <p>
                Lorem ipsum dolor sit, amet consectetur adipisicing elit.
                Doloribus, quasi.
              </p>
            </div>
          </div>
          <div className="w-full text-white flex items-start gap-2">
            <p className="text-5xl">2</p>
            <div>
              <p className="mb-2 text-lg font-medium">Insight & Analytics</p>
              <p>
                Lorem ipsum dolor sit, amet consectetur adipisicing elit.
                Doloribus, quasi.
              </p>
            </div>
          </div>
          <div className="w-full text-white flex items-start gap-2">
            <p className="text-5xl">3</p>
            <div>
              <p className="mb-2 text-lg font-medium">Bio Links</p>
              <p>
                Lorem ipsum dolor sit, amet consectetur adipisicing elit.
                Doloribus, quasi.
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

export default HomepageAbout;
