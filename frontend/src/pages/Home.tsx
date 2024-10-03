import { Button, Carousel } from "flowbite-react";
import { buttonCustom } from "../components/Button";

export default function Home() {
  return (
    //Hero
    <div className="grid grid-cols-2 grid-rows-3 max-h-[30rem] gap-x-10 mt-10">
      <h1 className="text-6xl text-start dark:text-white self-end">
        Can Learning Be Fun? Our Games Prove it!
      </h1>
      <h3 className="col-start-1 text-start self-center row-start-2 text-2xl dark:text-white">
        Explore innovative indie games that blend fun with learning or dive into
        immersive gameplay for pure enjoyment.
      </h3>
      <Button
        theme={buttonCustom}
        color="primary"
        className="max-w-60 self-center"
      >
        Explore Our Games
      </Button>
      <div className="h-56 sm:h-64 xl:h-96 2xl:h-[30rem] col-start-2 row-start-1 row-span-3">
        <Carousel slideInterval={5000}>
          <img
            src="https://flowbite.com/docs/images/carousel/carousel-1.svg"
            alt="..."
          />
          <img
            src="https://flowbite.com/docs/images/carousel/carousel-2.svg"
            alt="..."
          />
          <img
            src="https://flowbite.com/docs/images/carousel/carousel-3.svg"
            alt="..."
          />
          <img
            src="https://flowbite.com/docs/images/carousel/carousel-4.svg"
            alt="..."
          />
          <img
            src="https://flowbite.com/docs/images/carousel/carousel-5.svg"
            alt="..."
          />
        </Carousel>
      </div>
    </div>
  );
}
