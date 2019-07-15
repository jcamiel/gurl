package query

import (
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	html1 = `<!DOCTYPE html><html lang="en-US">
<head>
<title>Hello,World!</title>
</head>
<body>
<div class="container">
<header>
	<!-- Logo -->
   <h1>City Gallery</h1>
</header>  
<nav>
  <ul>
    <li><a href="/London">London</a></li>
    <li><a href="/Paris">Paris</a></li>
    <li><a href="/Tokyo">Tokyo</a></li>
  </ul>
</nav>
<article>
  <h1>London</h1>
  <img src="pic_mountain.jpg" alt="Mountain View" style="width:304px;height:228px;">
  <p>London is the capital city of England. It is the most populous city in the  United Kingdom, with a metropolitan area of over 13 million inhabitants.</p>
  <p>Standing on the River Thames, London has been a major settlement for two millennia, its history going back to its founding by the Romans, who named it Londinium.</p>
</article>
<footer>Copyright &copy; W3Schools.com</footer>
</div>
</body>
</html>
`
	html2 = `
<!DOCTYPE HTML>
<html lang="fr"
      xmlns="http://www.w3.org/1999/xhtml">






<head>
    <title>Offres internet Livebox Fibre dès 57,40 €/mois - Orange Réunion</title>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0">

    <link rel="icon" type="image/png" href="http://localhost:8080/images/logo-orange-32.png">

    

        <meta name="title" content="Offres internet Livebox Fibre dès 57,40 €/mois - Orange Réunion">
        <meta name="description" content="Testez votre éligibilité à la fibre optique Orange à la Réunion, internet très haut débit avec Livebox 4 + nouveau décodeur TV UHD inclus.">
        <meta data-name="version" content="1.11.0-SNAPSHOT">
        <meta data-name="git-ref" content="db786786">
        <meta data-name="tms-url" content="//tags.tiqcdn.com/utag/orange/caraibe/qa/utag.js">

        

    

    
        <meta data-name="datalayer" content="{&quot;type_page&quot;:&quot;liste_offre&quot;,&quot;canal&quot;:&quot;web&quot;,&quot;segment&quot;:&quot;REU&quot;,&quot;couleur&quot;:&quot;orange&quot;,&quot;titre_page&quot;:&quot;Offres internet Livebox Fibre dès 57,40 €/mois - Orange Réunion&quot;,&quot;domaine&quot;:&quot;shop&quot;,&quot;univers_affichage&quot;:&quot;internet&quot;,&quot;sous_univers&quot;:&quot;fibre&quot;}">
    

    
        <meta name="_csrf_token" content="0fadaca4-0f21-42c0-98d8-2c271ea1b1cf"/>
        <meta name="_csrf_header" content="X-CSRF-TOKEN"/>
    

    <script src="http://localhost:8080/js/vendor.js" type="text/javascript"></script>

<link rel="stylesheet" type="text/css" href="http://localhost:8080/css/reunion.css">

</head>

<body>

<div role="main">

    <div class="limit-max-width">
        <div class="container-fluid">

            <div class="d-none d-lg-block">
                <div class="row pt-2 mx-5">
                    <div>

    <div class="col-12">

        <div class="d-lg-none">
            <a href="https://reunion.orange.fr/boutique/univers-internet/"
               class="o-link-arrow back tms-listener"
               data-tagzone="fil_d_ariane"
               data-tagnom="retour"
               data-tagcible="https://reunion.orange.fr/boutique/univers-internet/">Offres Fibre</a>
        </div>

        <div class="d-none d-lg-block">
            <nav aria-label="breadcrumb">
                <ol class="breadcrumb bg-transparent">
                    
                        

    
        <li class="breadcrumb-item">
            <a class="tms-listener"
               href="https://reunion.orange.fr/"
               data-tagzone="fil_d_ariane"
               data-tagnom="Accueil"
               data-tagcible="https://reunion.orange.fr/">Accueil</a>
            
        </li>
    


                    
                        

    
        <li class="breadcrumb-item">
            <a class="tms-listener"
               href="https://reunion.orange.fr/boutique/"
               data-tagzone="fil_d_ariane"
               data-tagnom="Boutique"
               data-tagcible="https://reunion.orange.fr/boutique/">Boutique</a>
            
        </li>
    


                    
                        

    
        <li class="breadcrumb-item">
            <a class="tms-listener"
               href="https://reunion.orange.fr/boutique/univers-internet/"
               data-tagzone="fil_d_ariane"
               data-tagnom="Univers Internet"
               data-tagcible="https://reunion.orange.fr/boutique/univers-internet/">Univers Internet</a>
            
        </li>
    


                    
                    <li class="breadcrumb-item active" aria-current="page">Offres Fibre</li>
                </ol>

            </nav>
        </div>

    </div>

</div>
                </div>
            </div>

            <div class="row py-2 px-sm-3 px-md-2 pt-lg-0">
                <div class="col-12 bg-white mx-lg-3 px-lg-3">
                    <div class="mb-0 mt-2 home-title-resp mt-lg-0">
                        <h1>Bienvenue sur la boutique Fibre de La Réunion</h1>
                    </div>
                    <div class="py-2 py-md-4 mb-md-2 home-subtitle-resp"><h2>Passez à la Fibre Orange</h2></div>
                </div>
            </div>

        </div>
        <div class="d-md-none">
            <hr class="py-0 mt-2 mb-3">
        </div>

        <div><div class="margin-0 offers">

    <div class="mx-lg-5 px-lg-5">
        <div class="row px-0 mx-0 px-lg-4 mx-lg-4">

            <div class="offerxyz">
                <div class="row py-2 pl-4 pr-3 pr-md-1 mx-2">
                    <div class="col-12 px-0">

                        <div class="offer-promo text-primary mt-2 pt-2">
                            <strong>Promo</strong>
                        </div>

                        <div class="offer-title font-size-18 mb-2 pb-2"><h2
                                aria-label="livebox magique"><strong>Livebox Magik Fibre 300Mbit/s</strong>
                        </h2>
                        </div>

                        <div>
    

    <div class="offer-price"><div class="recurring-price">

    <div class="left-part text-primary">
        <div class="font-weight-bold">47</div>
    </div>

    <div class="right-part">
        <div class="cents-part text-primary"><strong>,40 €</strong></div>
        
            <div class="font-weight-bold">/<span class="recurring-part">mois</span></div>
        
    </div>
</div></div>
    <div class="font-size-14 line-height-20">
        <div class="textGrey">pendant 12 mois puis <strong>57,40 €/mois</strong></div>
        <div class="textGrey">Location Livebox incluse</div>
        <div class="pb-2"><strong>Engagement 12 mois</strong></div>
    </div>
</div>

                    </div>
                </div>

                <div class="row pb-2 px-4 mx-2">
                    <div class="col-12 px-0 button">
                        <a class="btn btn-info tms-listener" href="http://localhost:8080/eligibilite?offre=magik-fibre-300"
                           data-tagnom="selectionner"
                           data-tagzone="livebox_magik_fibre_300mbits"
                           data-tagcible="http://localhost:8080/eligibilite?offre=magik-fibre-300">Sélectionner</a>
                    </div>
                </div>

                <div>
                    <div>
    
    <div class="offer-attribute row py-2 px-4 mx-2">
        <div class="col-1 px-0 mt-1">
            <img src="http://localhost:8080/images/pictos/picto-internet.png"
                 class="img-fluid" alt="picto world wide web" srcset="http://localhost:8080/images/pictos/picto-internet@2x.png 2x, http://localhost:8080/images/pictos/picto-internet@3x.png 3x">
        </div>
        <div class="col-11 px-sm-0 pl-2 pl-md-2">
                             <span class="font-size-14 textGrey">Internet très haut débit jusqu&apos;à <strong>300Mbits/s</strong> en téléchargement et <strong>300 Mbit/s</strong> en envoi</span>
        </div>
    </div>
</div>
                </div>

                <div>
                    <div>
    
    <div class="offer-attribute row py-2 px-4 mx-2">
        <div class="col-1 px-0 mt-1">
            <img src="http://localhost:8080/images/pictos/picto-tv.png"
                 class="img-fluid" alt="picto world wide web" srcset="http://localhost:8080/images/pictos/picto-tv@2x.png 2x, http://localhost:8080/images/pictos/picto-tv@3x.png 3x">
        </div>
        <div class="col-11 px-sm-0 pl-2 pl-md-2">
                             <span class="font-size-14 textGrey">TV d&apos;Orange : jusqu&apos;à 110 chaînes et services inclus <strong>+ OCS offert</strong></span>
        </div>
    </div>
</div>
                </div>

                <div>
                    <div>
    
    <div class="offer-attribute row py-2 px-4 mx-2">
        <div class="col-1 px-0 mt-1">
            <img src="http://localhost:8080/images/pictos/picto-appels.png"
                 class="img-fluid" alt="picto world wide web" srcset="http://localhost:8080/images/pictos/picto-appels@2x.png 2x, http://localhost:8080/images/pictos/picto-appels@3x.png 3x">
        </div>
        <div class="col-11 px-sm-0 pl-2 pl-md-2">
                             <span class="font-size-14 textGrey">Illimités vers les fixes de 100 destinations et <strong>les mobiles de La Réunion, Mayotte, autres DOM + Métropole</strong></span>
        </div>
    </div>
</div>
                </div>

                <div>
                    <div>
    
    <div class="offer-attribute row py-2 px-4 mx-2">
        <div class="col-1 px-0 mt-1">
            <img src="http://localhost:8080/images/pictos/picto-modem.png"
                 class="img-fluid" alt="picto world wide web" srcset="http://localhost:8080/images/pictos/picto-modem@2x.png 2x, http://localhost:8080/images/pictos/picto-modem@3x.png 3x">
        </div>
        <div class="col-11 px-sm-0 pl-2 pl-md-2">
                             <span class="font-size-14 textGrey">Livebox 4<br/>Nouveau Décodeur TV 4</span>
        </div>
    </div>
</div>
                </div>

                
                    <div class="offer-odr bg-light py-3 px-4 mt-4" height="59px">
                        <div><div>
    

    <div class="media px-2">
        <img src="http://localhost:8080/images/pictos/picto-cadeau.png"
             class="img-fluid m-sm-0 mt-1" alt="picto cadeau avec 100€ remboursés en offre de bienvenue" srcset="http://localhost:8080/images/pictos/picto-cadeau@2x.png 2x, http://localhost:8080/images/pictos/picto-cadeau@3x.png 3x">

        <div class="media-body align-self-center px-2 font-size-14">
            <strong>100€ remboursés en cadeau de bienvenue.</strong>
            <a class="py-3 tms-listener"
               href="https://reunion.orange.fr/offres-de-remboursement/internet/"
               target="_blank"
               data-tagzone="livebox_magik_fibre_300mbits"
               data-tagnom="odr"
               data-tagcible="https://reunion.orange.fr/offres-de-remboursement/internet/">
                <div class="px-0 pt-1">
                        <strong class="underline">Voir conditions</strong>
                        <strong class="icon icon-arrow-next xsmall" aria-hidden="true"></strong>
                </div>
            </a>
        </div>
    </div>
</div></div>
                    </div>
                

                <div class="row mb-2 py-4 px-4 mx-2 mt-2" style="display: none;">
                    <div class="col-12 px-0 button">
                        <a class="btn btn-secondary tms-listener" href="#"
                           data-tagnom="voir_le_detail"
                           data-tagzone="livebox_magik_fibre_300mbits"
                           data-tagcible="TBD">Voir le détail</a>
                    </div>

                </div>
                <div class="py-3" style="display: block;">
                </div>

                <div class="d-md-none limit-div"></div>

            </div>

            <div class="offerxyz">
                <div class="row py-2 pl-4 pr-3 pr-md-1 mx-2">
                    <div class="col-12 px-0">

                        <div class="offer-promo text-primary mt-2 pt-2">
                            <strong>Promo</strong>
                        </div>

                        <div class="offer-title font-size-18 mb-2 pb-2"><h2
                                aria-label="livebox magique"><strong> Livebox Magik Fibre 1Gbit/s </strong>
                        </h2>
                        </div>

                        <div>
    

    <div class="offer-price"><div class="recurring-price">

    <div class="left-part text-primary">
        <div class="font-weight-bold">57</div>
    </div>

    <div class="right-part">
        <div class="cents-part text-primary"><strong>,40 €</strong></div>
        
            <div class="font-weight-bold">/<span class="recurring-part">mois</span></div>
        
    </div>
</div></div>
    <div class="font-size-14 line-height-20">
        <div class="textGrey">pendant 12 mois puis <strong>67,40 €/mois</strong></div>
        <div class="textGrey">Location Livebox incluse</div>
        <div class="pb-2"><strong>Engagement 12 mois</strong></div>
    </div>
</div>

                    </div>
                </div>

                <div class="row pb-2 px-4 mx-2">
                    <div class="col-12 px-0 button">
                        <a class="btn btn-info tms-listener" href="http://localhost:8080/eligibilite?offre=magik-fibre-1000"
                           data-tagnom="selectionner"
                           data-tagzone="_livebox_magik_fibre_1gbits_"
                           data-tagcible="http://localhost:8080/eligibilite?offre=magik-fibre-1000">Sélectionner</a>
                    </div>
                </div>

                <div>
                    <div>
    
    <div class="offer-attribute row py-2 px-4 mx-2">
        <div class="col-1 px-0 mt-1">
            <img src="http://localhost:8080/images/pictos/picto-internet.png"
                 class="img-fluid" alt="picto world wide web" srcset="http://localhost:8080/images/pictos/picto-internet@2x.png 2x, http://localhost:8080/images/pictos/picto-internet@3x.png 3x">
        </div>
        <div class="col-11 px-sm-0 pl-2 pl-md-2">
                             <span class="font-size-14 textGrey">Internet très haut débit jusqu&apos;à <strong>1 Gbit/s</strong> en téléchargement et <strong>300 Mbit/s</strong> en envoi</span>
        </div>
    </div>
</div>
                </div>

                <div>
                    <div>
    
    <div class="offer-attribute row py-2 px-4 mx-2">
        <div class="col-1 px-0 mt-1">
            <img src="http://localhost:8080/images/pictos/picto-tv.png"
                 class="img-fluid" alt="picto world wide web" srcset="http://localhost:8080/images/pictos/picto-tv@2x.png 2x, http://localhost:8080/images/pictos/picto-tv@3x.png 3x">
        </div>
        <div class="col-11 px-sm-0 pl-2 pl-md-2">
                             <span class="font-size-14 textGrey">TV d&apos;Orange : jusqu&apos;à 110 chaînes et services inclus <strong>+ OCS offert</strong></span>
        </div>
    </div>
</div>
                </div>

                <div>
                    <div>
    
    <div class="offer-attribute row py-2 px-4 mx-2">
        <div class="col-1 px-0 mt-1">
            <img src="http://localhost:8080/images/pictos/picto-appels.png"
                 class="img-fluid" alt="picto world wide web" srcset="http://localhost:8080/images/pictos/picto-appels@2x.png 2x, http://localhost:8080/images/pictos/picto-appels@3x.png 3x">
        </div>
        <div class="col-11 px-sm-0 pl-2 pl-md-2">
                             <span class="font-size-14 textGrey">Illimités vers les fixes de 100 destinations et <strong>les mobiles de La Réunion, Mayotte, autres DOM + Métropole</strong></span>
        </div>
    </div>
</div>
                </div>

                <div>
                    <div>
    
    <div class="offer-attribute row py-2 px-4 mx-2">
        <div class="col-1 px-0 mt-1">
            <img src="http://localhost:8080/images/pictos/picto-modem.png"
                 class="img-fluid" alt="picto world wide web" srcset="http://localhost:8080/images/pictos/picto-modem@2x.png 2x, http://localhost:8080/images/pictos/picto-modem@3x.png 3x">
        </div>
        <div class="col-11 px-sm-0 pl-2 pl-md-2">
                             <span class="font-size-14 textGrey">Livebox 4<br/>Nouveau Décodeur TV 4</span>
        </div>
    </div>
</div>
                </div>

                
                    <div class="offer-odr bg-light py-3 px-4 mt-4" height="59px">
                        <div><div>
    

    <div class="media px-2">
        <img src="http://localhost:8080/images/pictos/picto-cadeau.png"
             class="img-fluid m-sm-0 mt-1" alt="picto cadeau avec 100€ remboursés en offre de bienvenue" srcset="http://localhost:8080/images/pictos/picto-cadeau@2x.png 2x, http://localhost:8080/images/pictos/picto-cadeau@3x.png 3x">

        <div class="media-body align-self-center px-2 font-size-14">
            <strong>100€ remboursés en cadeau de bienvenue.</strong>
            <a class="py-3 tms-listener"
               href="https://reunion.orange.fr/offres-de-remboursement/internet/"
               target="_blank"
               data-tagzone="_livebox_magik_fibre_1gbits_"
               data-tagnom="odr"
               data-tagcible="https://reunion.orange.fr/offres-de-remboursement/internet/">
                <div class="px-0 pt-1">
                        <strong class="underline">Voir conditions</strong>
                        <strong class="icon icon-arrow-next xsmall" aria-hidden="true"></strong>
                </div>
            </a>
        </div>
    </div>
</div></div>
                    </div>
                

                <div class="row mb-2 py-4 px-4 mx-2 mt-2" style="display: none;">
                    <div class="col-12 px-0 button">
                        <a class="btn btn-secondary tms-listener" href="#"
                           data-tagnom="voir_le_detail"
                           data-tagzone="_livebox_magik_fibre_1gbits_"
                           data-tagcible="TBD">Voir le détail</a>
                    </div>

                </div>
                <div class="py-3" style="display: block;">
                </div>

                

            </div>
        </div>
    </div>
</div></div>

        <div><div class="advantage">

    <div class="d-md-none">
        <div><div class="margin-0 advantage-div">

        <div class="row p-0 m-0">
            <div class="py-2 col-12 px-0">
                <div class="row m-0 p-0 pt-2">

                    <div class="col-12 col-md-6 px-0 pt-3">

                        <div class="advantage-img">
                            <div><span class="responsive-image">

    <picture>
         <source media="(max-width: 479px)"
                 srcset="http://localhost:8080/images/pages/home/advantage/320-Live-Box.png 1x, http://localhost:8080/images/pages/home/advantage/320-Live-Box@2x.png 2x, http://localhost:8080/images/pages/home/advantage/320-Live-Box@3x.png 3x"/>
        <source media="(min-width: 480px) and (max-width: 767px)"
                srcset="http://localhost:8080/images/pages/home/advantage/480-Live-Box.png 1x, http://localhost:8080/images/pages/home/advantage/480-Live-Box@2x.png 2x, http://localhost:8080/images/pages/home/advantage/480-Live-Box@3x.png 3x"/>
        <source media="(min-width: 768px) and (max-width: 1023px)"
                srcset="http://localhost:8080/images/pages/home/advantage/768-Live-Box.png 1x, http://localhost:8080/images/pages/home/advantage/768-Live-Box@2x.png 2x, http://localhost:8080/images/pages/home/advantage/768-Live-Box@3x.png 3x"/>
        <source media="(min-width: 1024px)"
                srcset="http://localhost:8080/images/pages/home/advantage/1024-Live-Box.png 1x, http://localhost:8080/images/pages/home/advantage/1024-Live-Box@2x.png 2x, http://localhost:8080/images/pages/home/advantage/1024-Live-Box@3x.png 3x"/>
        <img alt="" class="img-fluid img-full-width" src="http://localhost:8080/images/pages/home/advantage/1024-Live-Box.png">
    </picture>

</span></div>
                        </div>

                        <div class="advantage-img" style="padding-top:55px">
                            <div><span class="responsive-image">

    <picture class="px-2">
         <source media="(max-width: 479px)"
                 srcset="http://localhost:8080/images/pages/home/advantage/320-Plus.png 1x, http://localhost:8080/images/pages/home/advantage/320-Plus@2x.png 2x, http://localhost:8080/images/pages/home/advantage/320-Plus@3x.png 3x"/>
        <source media="(min-width: 480px) and (max-width: 767px)"
                srcset="http://localhost:8080/images/pages/home/advantage/480-Plus.png 1x, http://localhost:8080/images/pages/home/advantage/480-Plus@2x.png 2x, http://localhost:8080/images/pages/home/advantage/480-Plus@3x.png 3x"/>
        <source media="(min-width: 768px) and (max-width: 1023px)"
                srcset="http://localhost:8080/images/pages/home/advantage/768-Plus.png 1x, http://localhost:8080/images/pages/home/advantage/768-Plus@2x.png 2x, http://localhost:8080/images/pages/home/advantage/768-Plus@3x.png 3x"/>
        <source media="(min-width: 1024px)"
                srcset="http://localhost:8080/images/pages/home/advantage/1024-Plus.png 1x, http://localhost:8080/images/pages/home/advantage/1024-Plus@2x.png 2x, http://localhost:8080/images/pages/home/advantage/1024-Plus@3x.png 3x"/>
        <img alt="" class="img-fluid img-full-width" src="http://localhost:8080/images/pages/home/advantage/1024-Plus.png">
    </picture>

</span></div>
                        </div>

                        <div class="advantage-img" style="padding-top:45px">
                            <div><span class="responsive-image">

    <picture>
         <source media="(max-width: 479px)"
                 srcset="http://localhost:8080/images/pages/home/advantage/320-SIM.png 1x, http://localhost:8080/images/pages/home/advantage/320-SIM@2x.png 2x, http://localhost:8080/images/pages/home/advantage/320-SIM@3x.png 3x"/>
        <source media="(min-width: 480px) and (max-width: 767px)"
                srcset="http://localhost:8080/images/pages/home/advantage/480-SIM.png 1x, http://localhost:8080/images/pages/home/advantage/480-SIM@2x.png 2x, http://localhost:8080/images/pages/home/advantage/480-SIM@3x.png 3x"/>
        <source media="(min-width: 768px) and (max-width: 1023px)"
                srcset="http://localhost:8080/images/pages/home/advantage/768-SIM.png 1x, http://localhost:8080/images/pages/home/advantage/768-SIM@2x.png 2x, http://localhost:8080/images/pages/home/advantage/768-SIM@3x.png 3x"/>
        <source media="(min-width: 1024px)"
                srcset="http://localhost:8080/images/pages/home/advantage/1024-SIM.png 1x, http://localhost:8080/images/pages/home/advantage/1024-SIM@2x.png 2x, http://localhost:8080/images/pages/home/advantage/1024-SIM@3x.png 3x"/>
        <img alt="" class="img-fluid img-full-width" src="http://localhost:8080/images/pages/home/advantage/1024-SIM.png">
    </picture>

</span></div>
                        </div>
                    </div>

                    <div><div>
    <div class="col-12 col-md-6 p-0">
        <div class="advantage-text-button">
            <div class="row m-0 px-4">
                <div class="col-12 pb-2">
                    <span>Votre avantage : <span class="text-primary">2 € à 10 €</span> de remise chaque mois sur votre forfait mobile.</span>
                </div>
                <div class="col-12 px-0 pt-2">
                    <div class="row m-0 px-0">
                        <div class="col-12 button">
                            <a href="#" data-toggle="modal" data-target="#advantage-popin"
                               class="btn btn-inverse btn-info tms-listener"
                               data-tagzone="bloc_avantages"
                               data-tagnom="comment_en_profiter"
                               data-tagcible="vers_popin_avantages">Comment en profiter</a>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </div>
</div></div>

                </div>
            </div>
        </div>


</div></div>
    </div>
    <div class="d-none d-md-block">
        <div><div class="margin-0 advantage-div">

        <div class="row py-2 mt-4">
            <div class="py-2 col-12">
                <div class="row m-0 p-0">

                    <div class="col-12 col-md-6 px-0 pt-3">

                        <div class="advantage-img pl-4">
                             <div><span class="responsive-image">

    <picture class="pl-lg-5 ml-lg-5">
         <source media="(max-width: 479px)"
                 srcset="http://localhost:8080/images/pages/home/advantage/320-Live-Box.png 1x, http://localhost:8080/images/pages/home/advantage/320-Live-Box@2x.png 2x, http://localhost:8080/images/pages/home/advantage/320-Live-Box@3x.png 3x"/>
        <source media="(min-width: 480px) and (max-width: 767px)"
                srcset="http://localhost:8080/images/pages/home/advantage/480-Live-Box.png 1x, http://localhost:8080/images/pages/home/advantage/480-Live-Box@2x.png 2x, http://localhost:8080/images/pages/home/advantage/480-Live-Box@3x.png 3x"/>
        <source media="(min-width: 768px) and (max-width: 1023px)"
                srcset="http://localhost:8080/images/pages/home/advantage/768-Live-Box.png 1x, http://localhost:8080/images/pages/home/advantage/768-Live-Box@2x.png 2x, http://localhost:8080/images/pages/home/advantage/768-Live-Box@3x.png 3x"/>
        <source media="(min-width: 1024px)"
                srcset="http://localhost:8080/images/pages/home/advantage/1024-Live-Box.png 1x, http://localhost:8080/images/pages/home/advantage/1024-Live-Box@2x.png 2x, http://localhost:8080/images/pages/home/advantage/1024-Live-Box@3x.png 3x"/>
        <img alt="" class="img-fluid img-full-width" src="http://localhost:8080/images/pages/home/advantage/1024-Live-Box.png">
    </picture>

</span></div>
                        </div>

                        <div class="advantage-img" style="padding-top:68px">
                           <div><span class="responsive-image">

    <picture class="px-2">
         <source media="(max-width: 479px)"
                 srcset="http://localhost:8080/images/pages/home/advantage/320-Plus.png 1x, http://localhost:8080/images/pages/home/advantage/320-Plus@2x.png 2x, http://localhost:8080/images/pages/home/advantage/320-Plus@3x.png 3x"/>
        <source media="(min-width: 480px) and (max-width: 767px)"
                srcset="http://localhost:8080/images/pages/home/advantage/480-Plus.png 1x, http://localhost:8080/images/pages/home/advantage/480-Plus@2x.png 2x, http://localhost:8080/images/pages/home/advantage/480-Plus@3x.png 3x"/>
        <source media="(min-width: 768px) and (max-width: 1023px)"
                srcset="http://localhost:8080/images/pages/home/advantage/768-Plus.png 1x, http://localhost:8080/images/pages/home/advantage/768-Plus@2x.png 2x, http://localhost:8080/images/pages/home/advantage/768-Plus@3x.png 3x"/>
        <source media="(min-width: 1024px)"
                srcset="http://localhost:8080/images/pages/home/advantage/1024-Plus.png 1x, http://localhost:8080/images/pages/home/advantage/1024-Plus@2x.png 2x, http://localhost:8080/images/pages/home/advantage/1024-Plus@3x.png 3x"/>
        <img alt="" class="img-fluid img-full-width" src="http://localhost:8080/images/pages/home/advantage/1024-Plus.png">
    </picture>

</span></div>
                        </div>

                        <div class="advantage-img" style="padding-top:50px">
                            <div><span class="responsive-image">

    <picture>
         <source media="(max-width: 479px)"
                 srcset="http://localhost:8080/images/pages/home/advantage/320-SIM.png 1x, http://localhost:8080/images/pages/home/advantage/320-SIM@2x.png 2x, http://localhost:8080/images/pages/home/advantage/320-SIM@3x.png 3x"/>
        <source media="(min-width: 480px) and (max-width: 767px)"
                srcset="http://localhost:8080/images/pages/home/advantage/480-SIM.png 1x, http://localhost:8080/images/pages/home/advantage/480-SIM@2x.png 2x, http://localhost:8080/images/pages/home/advantage/480-SIM@3x.png 3x"/>
        <source media="(min-width: 768px) and (max-width: 1023px)"
                srcset="http://localhost:8080/images/pages/home/advantage/768-SIM.png 1x, http://localhost:8080/images/pages/home/advantage/768-SIM@2x.png 2x, http://localhost:8080/images/pages/home/advantage/768-SIM@3x.png 3x"/>
        <source media="(min-width: 1024px)"
                srcset="http://localhost:8080/images/pages/home/advantage/1024-SIM.png 1x, http://localhost:8080/images/pages/home/advantage/1024-SIM@2x.png 2x, http://localhost:8080/images/pages/home/advantage/1024-SIM@3x.png 3x"/>
        <img alt="" class="img-fluid img-full-width" src="http://localhost:8080/images/pages/home/advantage/1024-SIM.png">
    </picture>

</span></div>
                        </div>

                    </div>

                    
    <div class="col-12 col-md-6 p-0">
        <div class="advantage-text-button">
            <div class="row m-0 px-4">
                <div class="col-12 pb-2">
                    <span>Votre avantage : <span class="text-primary">2 € à 10 €</span> de remise chaque mois sur votre forfait mobile.</span>
                </div>
                <div class="col-12 px-0 pt-2">
                    <div class="row m-0 px-0">
                        <div class="col-12 button">
                            <a href="#" data-toggle="modal" data-target="#advantage-popin"
                               class="btn btn-inverse btn-info tms-listener"
                               data-tagzone="bloc_avantages"
                               data-tagnom="comment_en_profiter"
                               data-tagcible="vers_popin_avantages">Comment en profiter</a>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </div>


                </div>
            </div>
        </div>

</div></div>
    </div>

</div></div>

        <div class="advantage-popin">
            <div><div id="advantage-popin"
     class="modal fade modal-screen" role="dialog"
     aria-hidden="true" aria-labelledby="advantageTitle">

    <div class="modal-align-helper">
        <div class="modal-dialog" role="document">
            <div class="modal-content">

                 <div class="modal-header">
                    <button class="close tms-listener" type="button"
                            data-dismiss="modal"
                            aria-label="Close"
                            data-tagzone="pop_in_avantage"
                            data-tagnom="fermer_croix"
                            data-tagcible="fermer_croix">
                        <span aria-hidden="true">&times;</span>
                    </button>
                </div>

                <div class="modal-body flex-center center-horizontal">
                    <div class="container-fluid m-0 p-0">
                        <div class="row py-0 mx-0 bg-white">
                            <div class="col-12 text-left popin-advantage-margin">
                                <strong class="mb-2 font-size-26 line-height-30">Avantage 100% Orange ?</strong>
                            </div>
                            <div class="col-12 text-left mt-2 popin-advantage-margin">
                                <div class="mb-4 line-height-22">Choisir Orange pour le mobile et la Fibre, c&apos;est bénéficier de 2 € à 10 € de remise tous les mois sur son forfait mobile :</div>
                            </div>

                            <div class="mb-4 col-12 px-0 popin-advantage-div">
                                <div class="row px-0 pt-2 mx-0">
                                    <div class="col-6 popin-advantage-margin py-3">
                                        <div class="mb-2 font-size-14">
                                             <strong>Montant de votre forfait mobile ( options incluses )</strong>
                                        </div>
                                    </div>
                                    <div class="col-6 popin-advantage-margin pt-4 mt-1 pb-3">
                                        <div class="mb-2 font-size-14">
                                            <strong>Votre remise</strong>
                                        </div>
                                    </div>
                                </div>
                                <div>
    <div class="popin-advantage-hr"></div>
    <div class="px-0">
        <div class="row px-0 mx-0">
            <div class="col-6 popin-advantage-margin py-4">
                <div class="font-size-14">A partir de 40 €/mois</div>
            </div>
            <div class="col-6 popin-advantage-margin py-4">
                <div class="font-size-14">-10 € tous les mois</div>
            </div>
        </div>
        <div class="popin-advantage-hr"></div>
    </div>
    <div class="px-0">
        <div class="row px-0 mx-0">
            <div class="col-6 popin-advantage-margin py-4">
                <div class="font-size-14">Entre 25 € et 39,99 €/mois</div>
            </div>
            <div class="col-6 popin-advantage-margin py-4">
                <div class="font-size-14">-5 € tous les mois</div>
            </div>
        </div>
        <div class="popin-advantage-hr"></div>
    </div>
    <div class="px-0">
        <div class="row px-0 mx-0">
            <div class="col-6 popin-advantage-margin py-4">
                <div class="font-size-14">Moins de 25 €/mois</div>
            </div>
            <div class="col-6 popin-advantage-margin py-4">
                <div class="font-size-14">-2 € tous les mois</div>
            </div>
        </div>
        
    </div>

</div>
                            </div>

                                <div class="col-12 text-left popin-advantage-margin pt-3">
                                    <div class="mb-1 mb-sm-3 font-size-20 line-height-26">
                                        <strong>Comment en profiter ?</strong>
                                    </div>
                                    <div>

    
    

    <div class="mb-2 px-0">

        <div class="media m-0 pb-2">

            <img src="http://localhost:8080/images/pages/home/advantage/popin/picto-call-center.png"
                 class="img-fluid pt-3 pr-1" alt="picto world wide web" srcset="http://localhost:8080/images/pages/home/advantage/popin/picto-call-center@2x.png 2x, http://localhost:8080/images/pages/home/advantage/popin/picto-call-center@3x.png 3x">

            <div class="media-body align-self-center px-2 pt-2">
                
                <div class="font-size-14">
                    <strong>Le service clients</strong></div>
                <div class="font-size-14">Appelez gratuitement le 456 depuis votre mobile.</div>
            </div>
        </div>
    </div>

    
    

    <div class="mb-2 px-0">

        <div class="media m-0 pb-2">

            <img src="http://localhost:8080/images/pages/home/advantage/popin/picto-rendez-vous.png"
                 class="img-fluid pt-3 pr-1" alt="picto world wide web" srcset="http://localhost:8080/images/pages/home/advantage/popin/picto-rendez-vous@2x.png 2x, http://localhost:8080/images/pages/home/advantage/popin/picto-rendez-vous@3x.png 3x">

            <div class="media-body align-self-center px-2">
                <div>
                    <a class="font-size-14 tms-listener" href="https://reunion.orange.fr/assistance/nous-contacter/#trouver-une-boutique"
                       data-tagzone="pop_in_avantage"
                       data-tagnom="rdv_boutique"
                       data-tagcible="https://reunion.orange.fr/assistance/nous-contacter/#trouver-une-boutique">
                        <strong>Rendez-vous en boutique Orange</strong>
                    </a>
                </div>
                
                <div class="font-size-14"></div>
            </div>
        </div>
    </div>
</div>
                                </div>

                        </div>
                    </div>

                    <div class="row py-2">
                        <div class="col-3 col-sm-4 col-md-1 pt-4"></div>
                        <div class="col-6 col-sm-4 col-md-4 px-md-0">
                            <button class="btn btn-info btn-lg btn-block tms-listener" type="button"
                                    data-dismiss="modal"
                                    data-tagzone="pop_in_avantage"
                                    data-tagnom="fermer_bouton"
                                    data-tagcible="fermer_bouton">Fermer</button>
                        </div>
                        <div class="col-3 col-sm-4 col-md-7 pt-4"></div>
                    </div>
                </div>

            </div>
        </div>
    </div>
</div></div>
        </div>

        <div><div class="orange-infos">

    <div class="container-fluid px-0 py-0 pb-md-4 mb-md-2 mb-lg-4 ml-md-3 pl-lg-5">
        <div class="row pt-4 px-0 mx-0 mt-2">
            <div class="col-12">
                <div class="my-2 orange-infos-title-resp"><h2>Orange s&#39;occupe de tout</h2>
                </div>
            </div>
        </div>

        <div class="row m-0 py-4 px-2">
            <div class="col-12 col-md-6 pr-4">
                <div class="img-orange-infos">
                    <img width="100%" src="http://localhost:8080/images/components/orangeInfos/Conservons.jpg" alt="photo d’une femme appelant depuis le téléphone sur ligne fixe"
                         class="shadow-left"/>
                </div>
            </div>

            <div class="col-12 col-md-6 px-0 px-md-4 pt-4 pt-md-2 pb-4">
                <div class="m-0 orange-infos-subtitle-resp"><h2>Votre n° fixe conservé sans coupure (gratuit)</h2>
                </div>
                <div class="m-0 mt-md-4 pt-md-2 orange-infos-text-resp textGrey">Orange met en service votre ligne fixe sans coupure (durant votre changement d’opérateur), puis résilie le contrat associé à votre ligne fixe auprès de votre ancien opérateur.</div>
            </div>

        </div>

        <div class="row m-0 pb-4 px-2 px-md-4 pt-md-4 mt-md-4">
            <div class="col-12 col-md-6 py-4 px-4 d-none d-md-block">
                <div class="m-0 orange-infos-subtitle-resp"><h2>Votre raccordement Fibre (inclus)</h2>
                </div>
                <div class="m-0 mt-md-4 pt-md-2 orange-infos-text-resp textGrey">La Fibre Orange est installée jusqu&#39;à chez vous en 4 heures environ par un technicien, à la date de votre choix et dans le respect de votre intérieur.</div>
            </div>
            <div class="col-12 col-md-6 pr-4">
                <div class="img-orange-infos">
                    <img width="100%" src="http://localhost:8080/images/components/orangeInfos/Installation.jpg" alt="photo d’un technicien expert installant la fibre optique à domicile"
                         class="shadow-320-768"/>
                </div>

            </div>

            <div class="col-12 py-4 px-0 d-md-none">
                <div class="m-0 orange-infos-subtitle-resp">Votre raccordement Fibre (inclus)</div>
                <div class="m-0 orange-infos-text-resp textGrey">La Fibre Orange est installée jusqu&#39;à chez vous en 4 heures environ par un technicien, à la date de votre choix et dans le respect de votre intérieur.</div>
            </div>
        </div>
    </div>
</div></div>

    </div>

    <div><div class="reassurance bg-gray-f4f4f4">

    <div class="margin-0 limit-max-width">
        <div class="container-fluid margin-0">
            <div class="row px-0 pt-0 pb-5 pb-sm-5 pb-md-4 pb-lg-4 px-lg-4 mx-lg-3 m-0">
                <div class="col-12 px-0 px-md-2 px-lg-0 pt-3 pb-2">
                    <div class="font-size-resp-26 pt-md-3 pt-lg-3 pb-md-4 pb-lg-4"><h2>Pourquoi choisir la Fibre Orange ?</h2>
                    </div>
                </div>

                <div class="col-12 col-md-3 px-0 pt-2 pb-4 pb-md-5 pb-lg-5 px-md-4 px-lg-3 font-size-14 text-center">
                    <div>
    <div>
        <img alt="" src="http://localhost:8080/images/components/reassurance/qualite.png"  class="img-fluid" srcset="http://localhost:8080/images/components/reassurance/qualite@2x.png 2x, http://localhost:8080/images/components/reassurance/qualite@3x.png 3x">
    </div>

    <div class="d-sm-none d-md-block">
        <strong>La qualité du réseau</strong>
        <div>Une connexion internet fiable et puissante, même en cas d&apos;intempéries.</div>
    </div>
    <div class="d-none d-sm-block d-md-none width-290">
        <strong>La qualité du réseau</strong>
        <div>Une connexion internet fiable et puissante, même en cas d&apos;intempéries.</div>
    </div>
</div>
                </div>

                <div class="col-12 col-md-3 px-0 pt-2 pb-4 pb-md-5 pb-lg-5 px-md-4 px-lg-3 font-size-14 text-center">
                    <div>
    <div>
        <img alt="" src="http://localhost:8080/images/components/reassurance/experience.png"  class="img-fluid" srcset="http://localhost:8080/images/components/reassurance/experience@2x.png 2x, http://localhost:8080/images/components/reassurance/experience@3x.png 3x">
    </div>

    <div class="d-sm-none d-md-block">
        <strong>L&apos;expérience</strong>
        <div>Orange, leader sur la fibre optique, a déjà raccordé plus de 1500 communes en France dont 15 à la Réunion.</div>
    </div>
    <div class="d-none d-sm-block d-md-none width-290">
        <strong>L&apos;expérience</strong>
        <div>Orange, leader sur la fibre optique, a déjà raccordé plus de 1500 communes en France dont 15 à la Réunion.</div>
    </div>
</div>
                </div>

                <div class="col-12 col-md-3 px-0 pt-2 pb-4 pb-md-5 pb-lg-5 px-md-4 px-lg-3 font-size-14 text-center">
                    <div>
    <div>
        <img alt="" src="http://localhost:8080/images/components/reassurance/technique.png"  class="img-fluid" srcset="http://localhost:8080/images/components/reassurance/technique@2x.png 2x, http://localhost:8080/images/components/reassurance/technique@3x.png 3x">
    </div>

    <div class="d-sm-none d-md-block">
        <strong>Le service</strong>
        <div>Un technicien vient à votre domicile pour installer la Fibre et vous faire découvrir ses fonctionnalités.</div>
    </div>
    <div class="d-none d-sm-block d-md-none width-290">
        <strong>Le service</strong>
        <div>Un technicien vient à votre domicile pour installer la Fibre et vous faire découvrir ses fonctionnalités.</div>
    </div>
</div>
                </div>

                <div class="col-12 col-md-3 px-0 pt-2 pb-4 pb-md-5 pb-lg-5 px-md-4 px-lg-3 font-size-14 text-center">
                    <div>
    <div>
        <img alt="" src="http://localhost:8080/images/components/reassurance/paiement.png"  class="img-fluid" srcset="http://localhost:8080/images/components/reassurance/paiement@2x.png 2x, http://localhost:8080/images/components/reassurance/paiement@3x.png 3x">
    </div>

    <div class="d-sm-none d-md-block">
        <strong>Un paiement 100% sécurisé</strong>
        <div>Nous vous offrons le plus haut niveau de sécurité pour vos paiements en ligne.</div>
    </div>
    <div class="d-none d-sm-block d-md-none width-290">
        <strong>Un paiement 100% sécurisé</strong>
        <div>Nous vous offrons le plus haut niveau de sécurité pour vos paiements en ligne.</div>
    </div>
</div>
                </div>

            </div>
        </div>
    </div>
</div></div>

    <div id="fis"><div class="margin-0 limit-max-width">

    <div class="container-fluid">
        <div class="row py-2 px-0 px-md-2 px-lg-0 px-lg-4 mx-lg-3 m-0">
            <a class="py-3 tms-listener" href="https://documentscontractuels.orange.fr/les-offres-orange-internet_cfis_3321.pdf" target="_blank"
               data-tagzone="telecharger_la_fiche_dinformation_detaillee"
               data-tagnom="telecharger_fiche_info"
               data-tagcible="telecharger_fiche_info">
                <div class="col-12 px-0">
                    <div>
                        <strong>Télécharger la fiche d&apos;information détaillée</strong>
                        <span class="icon icon-arrow-next xsmall text-primary ml-2 ml-md-4" aria-hidden="true"></span>
                    </div>
                </div>
            </a>
        </div>
    </div>
</div></div>

    <div id="legalMentions"><div class="bg-gray-f4f4f4">

    <div class="margin-0 limit-max-width">
        <div class="container-fluid margin-0">
            <div class="px-md-2">
                <div class="row px-0 pt-0 pb-5 m-0 px-lg-4 mx-lg-1 line-height-18">

                    <div class="col-12 px-0 pt-4">
                        <strong class="font-size-16">Mentions légales</strong>
                    </div>

                    
                        <div>
    <div class="col-12 px-0 pt-3 pb-1 font-size-12 text-left">
        <strong>Offre</strong>
        <div>Offres valables à La Réunion à partir du 04/07/2019 sous réserve de compatibilité technique et d’éligibilité géographique de la ligne téléphonique du client aux différents services. Location Livebox 4 incluse à 2,50€/mois. Frais d’activation décodeur TV 50€.</div>
    </div>
</div>
                    
                        <div>
    <div class="col-12 px-0 pt-3 pb-1 font-size-12 text-left">
        <strong>Cadeau de bienvenue</strong>
        <div>Du 04/07/2019 au 09/10/2019 pour toute première souscription (hors changement d’offre) à une offre Livebox Magik Fibre avec engagement 12 mois. Remboursement sur facture suivant la réception complète des pièces justificatives et du formulaire de remboursement disponible sur orange.re, avant le 09/12/2019 inclus, le cachet de La Poste faisant foi. Un seul remboursement par souscription Internet. Toute demande illisible, incomplète ou ne répondant pas aux conditions de l’offre sera considérée comme non conforme et ne sera donc pas prise en compte.</div>
    </div>
</div>
                    
                        <div>
    <div class="col-12 px-0 pt-3 pb-1 font-size-12 text-left">
        <strong>Appels</strong>
        <div>Liste des destinations sur la fiche tarifaire en vigueur sur www.orange.re. Appels illimités 24h/24 hors numéros courts et spéciaux et jusqu’à 250 correspondants, différents/mois et 3h maximum/appel. Appels depuis le poste fixe branché sur la Livebox.</div>
    </div>
</div>
                    
                        <div>
    <div class="col-12 px-0 pt-3 pb-1 font-size-12 text-left">
        <strong>TV d&apos;Orange</strong>
        <div>Offres réservées aux abonnés Internet TV d’Orange, soumises à conditions, valables à La Réunion, sous réserve d’éligibilité technique, géographique, sous condition de versement de frais d’activation du décodeur TV de 50€ TTC. TV d’Orange sur support Fibre d’Orange. Voir conditions et détails sur orange.re. Chaînes accessibles sous réserve de l’accord des chaînes et du CSA. La liste des chaînes est susceptible d’évolution.</div>
    </div>
</div>
                    
                        <div>
    <div class="col-12 px-0 pt-3 pb-1 font-size-12 text-left">
        <strong>OCS</strong>
        <div>Service disponible sous réserve d’activation du service par le client et du versement de frais de mise en service de 50 Euros pour le décodeur TV.</div>
    </div>
</div>
                    
                        <div>
    <div class="col-12 px-0 pt-3 pb-1 font-size-12 text-left">
        <strong>Espace de stockage</strong>
        <div>Espace de stockage accessible depuis un équipement compatible sous réserve d’une connexion Internet. Pour un usage mobile ou tablette, les connexions à l’Internet mobile ne sont pas incluses dans le service et leur tarification varie selon l’offre mobile du client. Usages sur réseaux compatibles. Service accessible sur terminaux compatibles avec l’application le Cloud d’Orange sous réserve de téléchargement de l’application. Ne sont pas compris dans le service les contenus et services «payants».</div>
    </div>
</div>
                    
                        <div>
    <div class="col-12 px-0 pt-3 pb-1 font-size-12 text-left">
        <strong>Suite de sécurité</strong>
        <div>Service qui permet de bénéficier d&apos;un outil limitant les risques numériques, sous réserve d’activation du service. Service limité à 5 licences pour 5 équipements différents. Service fourni par un éditeur tiers. Avec équipement compatible : systèmes d’exploitation Microsoft Windows XP à 10 (32/64 bits), Mac (équipé d’un processeur Intel uniquement) 10.7 à 10.11 et Android 4 à 7.  Conditions techniques sur la page assistance du site orange.re</div>
    </div>
</div>
                    
                        <div>
    <div class="col-12 px-0 pt-3 pb-1 font-size-12 text-left">
        <strong>Décodeur TV</strong>
        <div>Décodeur TV 4 soumis aux conditions d’éligibilité et de compatibilité technique à la TV par internet. Disponible avec les offres Livebox Magik Fibre et télévision par internet uniquement. Livebox 4 nécessaire, location : 2,50€/mois.</div>
    </div>
</div>
                    
                        <div>
    <div class="col-12 px-0 pt-3 pb-1 font-size-12 text-left">
        <strong>Enregistreur TV</strong>
        <div>Capacité de stockage de l&apos;enregistreur jusqu&apos;à 450Go offert sur demande</div>
    </div>
</div>
                    
                        <div>
    <div class="col-12 px-0 pt-3 pb-1 font-size-12 text-left">
        <strong>Promotions</strong>
        <div>Offre valable à la Réunion du 04/07/2019 au 09/10/2019 pour toute souscription à une offre Livebox Magik Fibre, avec engagement de 12 mois. Cumulable avec les offres de remboursement en cours.</div>
    </div>
</div>
                    

                </div>
            </div>
        </div>
    </div>
</div></div>

</div>

<script src="http://localhost:8080/js/reunion.js" type="text/javascript"></script>

<a class="o-scroll-up" title="back to top">
    <span class="o-scroll-up-text d-none d-md-inline-block">Haut de page</span>
    <span class="o-scroll-up-icon" aria-hidden="true"></span>
</a>




    <script src="https://reunion.orange.fr/js/global-layout.js"></script>


</body>
</html>
`
	html3 = `<html>
<body>
<div class="add">
  <div>
    <span class="addone">
  </div>
</div>
<div class="add">
  <div>
    <span class="addone">
  </div>
</div>
</body>
</html>
`
)

func TestEvalXPathHTML(t *testing.T) {
	var tests = []struct {
		expr     string
		expected interface{}
	}{
		{`normalize-space(//div[@class="container"]/header)`, "City Gallery"},
		{`//li/a`, []string{"London", "Paris", "Tokyo"}},
		{`count(//h1)`, 2.0},
		{`boolean(count(//code))`, false},
	}
	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			v, err := EvalXPathHTML(test.expr, []byte(html1))
			assert.Equal(t, test.expected, v)
			assert.Nil(t, err)
		})
	}
}


func TestEvalXPathHTMLBug1(t *testing.T) {

	html := `<html>
	<body>
		<div class="fruit">Apple</div>
		<div class="fruit">Banana</div>
		<div class="fruit">Lemon</div>
	</body>
</html>`

	doc, _ := htmlquery.Parse(strings.NewReader(html))
	test := `string((//div[@class="fruit"])[2])`
	expr, _ := xpath.Compile(test)
	v := expr.Evaluate(htmlquery.CreateXPathNavigator(doc))
	assert.Equal(t, v, "Banana")
}

func TestEvalXPathHTMLBug2(t *testing.T) {

	html := `<html>
	<body>
			<div class="fruit">Apple</div>
			<div class="color">Red</div>
	</body>
</html>`

	doc, _ := htmlquery.Parse(strings.NewReader(html))
	test := `string((//div[@class="color"])[1])`
	expr, _ := xpath.Compile(test)
	v := expr.Evaluate(htmlquery.CreateXPathNavigator(doc))
	assert.Equal(t, v, "Red")
}