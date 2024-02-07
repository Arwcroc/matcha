import React from 'react';
import Box from '@mui/material/Box';
import BackImg from '/mnt/nfs/homes/tefroiss/Documents/matcha/matcha/src/Images/MainPage/backgroung-unsplash.jpg';
import Logo from '/mnt/nfs/homes/tefroiss/Documents/matcha/matcha/src/Images/MenuSection/uletter.png';
import Portrait from '/mnt/nfs/homes/tefroiss/Documents/matcha/matcha/src/Images/MainPage/joconde.jpg';
import { CarouselProvider, Slider, Slide, ButtonBack, ButtonNext } from 'pure-react-carousel';
import 'pure-react-carousel/dist/react-carousel.es.css';

const SlideBox = () => {
	return (
		<CarouselProvider
			naturalSlideWidth={100}
			naturalSlideHeight={100}
			totalSlides={5}
			className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Slider"
		>
			<ButtonBack className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Back">{"<"}</ButtonBack>
			<Slider className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Pics">
				<Slide index={0}><img src={Portrait} alt="Slide 1"/></Slide>
				<Slide index={1}><img src={Logo} alt="Slide 2" width="100%" height="100%"/></Slide>
				<Slide index={2}><img src={Portrait} alt="Slide 3" width="100%" height="100%"/></Slide>
				<Slide index={3}><img src={Logo} alt="Slide 4" width="100%" height="100%"/></Slide>
				<Slide index={4}><img src={Portrait} alt="Slide 5" width="100%" height="100%"/></Slide>
			</Slider>
			<ButtonNext className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Next">{">"}</ButtonNext>
		</CarouselProvider>
	);
}

const Swipe = () => {
	return (
		<div>
			<Box className="App__WebContainer__Website__Main">
				<img className="App__WebContainer__Website__Main__Background" src={BackImg} style={{ width:"100%"}}/>
				<Box className="App__WebContainer__Website__Main__OldOldCard">
					<img src={Portrait} className="App__WebContainer__Website__Main__OldOldCard__Profile"/>
				</Box>
				<Box className="App__WebContainer__Website__Main__OldCard">
					<img src={Portrait} className="App__WebContainer__Website__Main__OldCard__Profile"/>
				</Box>
				<Box className="App__WebContainer__Website__Main__PrimaryCard">
					<Box className="App__WebContainer__Website__Main__PrimaryCard__Profile">
						<SlideBox />
					</Box>
					<Box className="App__WebContainer__Website__Main__Response">
						<Box className="App__WebContainer__Website__Main__Response__No">
							Non
						</Box>
						<Box className="App__WebContainer__Website__Main__Response__Yes">
							Oui
						</Box>
					</Box>
				</Box>
				<Box className="App__WebContainer__Website__Main__NextCard">
					<img src={Logo} width={70} height={70} />
				</Box>
				<Box className="App__WebContainer__Website__Main__NextNextCard">
					<img src={Logo} width={50} height={50} opacity={0.2} />
				</Box>
			</Box>
		</div>
	);
}

export default Swipe;