package icecubeapp

const signedoutForm = `
<p>Thanks for visiting!</p>
<a href="/">Sign in again</a>
`

const sigCreateForm = stylesForm +
	`<fieldset class="fieldset-auto-width">
  <legend><h2>New Signal Creator</h2></legend>
  <ul id="menu">
  <li>` + readSigCreateForm + `</li>
  <li>` + writeSigCreateForm + `</li>
  </ul>
  </fieldset>
  <br>
  `

const readSigCreateForm = `
<form action="/createReadSig" method="post">
  <fieldset class="fieldset-auto-width">
    <legend><h2>Read</h2></legend>
    Signal name:<br>
    <input type="text" name="signame" required></input><br><br>
    Serial string:<br>
    <input type="text" name="serialcommand" required></input><br><br>` +
	sigSelectForm + "<br><br>" +
	`<input type="submit" value="Create">
    </fieldset>
</form>`

const writeSigCreateForm = `
<form action="/createWriteSig" method="post">
  <fieldset class="fieldset-auto-width">
    <legend><h2>Write</h2></legend>
    Signal name:<br>
    <input type="text" name="signame" required></input><br><br>
    Serial string:<br>
    <input type="text" name="serialcommand" required></input><br><br>` +
	sigSelectForm + "<br><br>" +
	`<input type="submit" value="Create">
    </fieldset>
</form>`

const sigSelectForm = `
  Data Type:<br>
  <select name="sigtype">
    <option value="integer">Integer</option>
    <option value="float">Float</option>
    <option value="string">C-style str</option>
  </select>`

const stylesForm = `
  <style type="text/css">
      .fieldset-auto-width {
           display: inline-block;
      }
      ul#menu li {
      display:inline-block;
      }
  </style>
  `
