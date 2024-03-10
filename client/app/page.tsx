import { HeaderComponent, FooterComponent } from "@/components/ui";
import {
  HomepageAbout,
  HomepageHero,
  HomepagePlatforms,
} from "@/components/home";

export default function Home() {
  return (
    <>
      <HeaderComponent />
      <main>
        <HomepageHero />
        <HomepagePlatforms />
        <HomepageAbout />
      </main>
      {/* <FooterComponent /> */}
    </>
  );
}
