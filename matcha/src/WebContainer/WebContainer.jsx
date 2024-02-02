import React from 'react';
// import './Website/Accueil.jsx'
import './Website/Swipe.jsx'
import './Header/MenuSection.jsx'
// import './Header/MenuSection_Connect.jsx'
import './Footer/Footer.jsx'
import MenuSection from './Header/MenuSection.jsx'
// import MenuSection_Connect from './Header/MenuSection_Connect.jsx'
// import Accueil from './Website/Accueil.jsx'
import Swipe from './Website/Swipe.jsx'
import Footer from './Footer/Footer.jsx'

const WebContainer = () => {
	return (
		<>
			<MenuSection/>
			<Swipe/>
			<Footer/>
		</>
	);
}

export default WebContainer;