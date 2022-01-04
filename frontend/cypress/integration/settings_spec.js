/// <reference types="cypress" />

import "../support/commands"

context('Settings Tests', () => {


  beforeEach(() => {
    cy.restoreLocalStorageCache();
  });

  afterEach(() => {
    cy.saveLocalStorageCache();
  });

  it('should Login', () => {
    cy.visit('/login')
    cy.get('#username').clear().type('admin')
    cy.get('#password').clear().type('admin')
    cy.get('button[type="submit"]').click()

    cy.get('.navbar-brand').should('contain', 'Statping')
    cy.getCookies()

    cy.getCookies().should('have.length', 1)
  })

  it('should confirm notifiers are installed', () => {
    cy.visit('/dashboard/settings')
    cy.get('#notifiers_tabs > a').should('have.length', 10)

    cy.get('#api_secret').should('not.have.value', '')
  })

  it('should update Statping settings', () => {
    cy.visit('/dashboard/settings')

    cy.get('#project').clear().type('Statping Updated')
    cy.get('#description').clear().type('Statping can use Cypress e2e testing to make it more stable!')
    cy.get('#domain').clear().type('http://localhost:8888')
    cy.get('#footer').clear().type('Statping Custom Footer')
    cy.get('#save_core').click()
  })

  it('should confirm Statping settings', () => {
    cy.visit('/dashboard/settings')

    cy.get('#project').should('have.value', 'Statping Updated')
    cy.get('#description').should('have.value', 'Statping can use Cypress e2e testing to make it more stable!')
    cy.get('#domain').should('have.value', 'http://localhost:8888')
    cy.get('#footer').should('have.value', 'Statping Custom Footer')
    cy.get('#api_secret').should('not.have.value', '')
  })

  it('should confirm new Footer text', () => {
    cy.visit('/dashboard/settings')
    cy.get('.footer').should('contain', 'Statping Custom Footer')
  })

  it('should regenerate API Keys', () => {
    cy.visit('/dashboard/settings')
    cy.get('#regenkeys').click()
    cy.get('#api_key').should('not.have.value', '')
    cy.get('#api_secret').should('not.have.value', '')
  })

  it('should create Local Assets', () => {
    cy.visit('/dashboard/settings')
    cy.get('#v-pills-style-tab').click()
    cy.get('#enable_assets').click()
    cy.wait(5000)
    cy.visit('/dashboard/settings')
    cy.get('#v-pills-style-tab').click()
    cy.get('#pills-vars-tab').click()
    cy.get('#assets_dir').should('contain', 'github.com/statping/statping/assets')
  })

  it('should save Local Assets', () => {
    cy.visit('/dashboard/settings')
    cy.get('#v-pills-style-tab').click()
    cy.get('#pills-vars-tab').click()
    cy.wait(1000)
    cy.get('.CodeMirror textarea').type('{downarrow}$example-variable: #bababa;{enter}', { force: true });
    cy.get('#save_assets').click();
  })

  it('should delete Local Assets', () => {
    cy.visit('/dashboard/settings')
    cy.get('#v-pills-style-tab').click()
    cy.get('#pills-vars-tab').click()
    cy.get('#delete_assets').click()
  })

  it('should view Cache', () => {
    cy.visit('/dashboard/settings')
    cy.get('#v-pills-cache-tab').click()
  })

})
