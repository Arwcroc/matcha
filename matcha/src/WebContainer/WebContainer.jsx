import React from 'react';
import './Website/Accueil.jsx'
import './Header/MenuSection.jsx'
import './Footer/Footer.jsx'
import MenuSection from './Header/MenuSection.jsx'
import Accueil from './Website/Accueil.jsx'
import Footer from './Footer/Footer.jsx'

const WebContainer = () => {
	return (
		<>
			<MenuSection/>
			<Accueil/>
			<Footer/>
		</>
	);
}

export default WebContainer;