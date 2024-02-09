import React from 'react';
import Box from '@mui/material/Box';
import BackImg from '../../Images/MainPage/backgroung-unsplash.jpg';
import Logo from '../../Images/MenuSection/uletter.png';
import Portrait from '../../Images/MainPage/joconde.jpg';
import GirlPort from '../../Images/MainPage/Girl_Pearl.jpg';
import Donna from '../../Images/MainPage/donna.jpg';
import Meuh from '../../Images/MainPage/meuhmeuh.jpg';
import Ermine from '../../Images/MainPage/Ermine.jpg';
import { CarouselProvider, Slider, Slide, ButtonBack, ButtonNext, DotGroup } from 'pure-react-carousel';
import 'pure-react-carousel/dist/react-carousel.es.css';

const SlideBox = () => {
	return (
		<CarouselProvider
			naturalSlideWidth={100}
			naturalSlideHeight={100}
			totalSlides={5}
			className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Slider"
			infinite
		>
			<DotGroup className="App__WebContainer__Website__Main__PrimaryCard__Carousel__DotGroup" />
			<ButtonBack className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Back">{"<"}</ButtonBack>
			<Slider className="App__WebContainer__Website__Main__PrimaryCard__Carousel__Pics">
				<Slide index={0}><div class="slide__image__container"><img class="slide__image__prout" src={Portrait} alt="Slide 1"/></div></Slide>
				<Slide index={1}><div class="slide__image__container"><img class="slide__image__prout" src={GirlPort} alt="Slide 2"/></div></Slide>
				<Slide index={2}><div class="slide__image__container"><img class="slide__image__prout" src={Meuh} alt="Slide 3"/></div></Slide>
				<Slide index={3}><div class="slide__image__container"><img class="slide__image__prout" src={Ermine} alt="Slide 4"/></div></Slide>
				<Slide index={4}><div class="slide__image__container"><img class="slide__image__prout" src={Donna} alt="Slide 5"/></div></Slide>
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