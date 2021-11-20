package webhookperso

/**
* @rename TokenPerso_exemple => TokenPerso
* @param string
* @return string
* @brief Renvoie le token discord
* @author Alexandre Caussades
 */

type Webhook_ struct {
	Token string
}

func TokenPerso_exemple() string {

	url := Webhook{
		Token: "", //Token discord
	}

	return url.Token

}
