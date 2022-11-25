package util

func FBTemplate() string {
	STR := `
	<section>
      <div class="container">
          <div class="info">
              <img src="https://raw.githubusercontent.com/hamzazaouya/Facebook-SignIn-LogIn-Page-Html-Css/master/img/facebook.svg" alt="Facebook" />
              <p>Facebook helps you connect and share with the people in your life.</p>
          </div>

          <div class="login">
              <form>
                  <input type="email" name="email" placeholder="Email address or phone number" />
                  <input type="password" name="password" placeholder="Password" />
                  <button>log in</button>
                  <a href="#">Forgotten account?</a>
                  <span></span>
                  <button>create new account</button>
              </form>
              <p><a href="#">Creat a Page</a> for a celebrity, band or business</p>
          </div>
      </div>
  </section>
  <footer>
      <div class="container">
          <div class="languages">
              <a href="index_en.html">English(US)</a>
              <a href="index_fr.html">Francais (France)</a>
              <a href="index_ar.html">العربية</a>
              <a href="index_amz.html">ⵜⴰⵎⴰⵣⵉⵖⵜ</a>
          </div>
          <div class="copyright">Facebook &copy; 2020</div>
      </div>
  </footer>`
	return STR
}
